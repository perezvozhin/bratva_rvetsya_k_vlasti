import { useCallback, useEffect, useRef, useState } from 'react'
import ChatList from './components/ChatList/ChatList'
import ChatWindow from './components/ChatWindow/ChatWindow'
import ConfirmDialog from './components/ConfirmDialog/ConfirmDialog'
import VacancyBrowser from './components/VacancyBrowser/VacancyBrowser'
import PromptDialog from './components/PromptDialog/PromptDialog'
import {
  getChats,
  getMessages,
  createChat,
  deleteChat,
  uploadInterview,
  streamMessage,
  getVacancies,
  searchVacancies,
  UnauthorizedError,
} from './api'
import { QUIZ_PROMPT, QUIZ_MARKER, extractQuiz } from './shared/quiz'
import s from './App.module.css'

const PER_PAGE = 20

// бэкенд отдаёт RFC3339 с наносекундами - в сайдбаре нужна короткая дата,
// иначе .time (flex-shrink: 0) выдавливает название чата из строки
function formatUpdatedAt(value) {
  if (!value) return ''

  const date = new Date(value)
  // уже готовая строка вроде 'сейчас' - отдаём как есть
  if (Number.isNaN(date.getTime())) return value

  const now = new Date()
  const sameDay =
    date.getDate() === now.getDate() &&
    date.getMonth() === now.getMonth() &&
    date.getFullYear() === now.getFullYear()

  return sameDay
    ? date.toLocaleTimeString(undefined, { hour: '2-digit', minute: '2-digit' })
    : date.toLocaleDateString(undefined, { day: '2-digit', month: '2-digit' })
}

function normalizeChat(raw) {
  const name = raw.chatName ?? raw.name ?? String(raw)

  return {
    name,
    company: raw.company ?? name,
    title: raw.title ?? '',
    lastMessage: raw.lastMessage ?? '',
    updatedAt: formatUpdatedAt(raw.updatedAt),
  }
}

function normalizeMessage(raw, index) {
  const text = raw.text ?? ''
  const isModel = raw.role === 'model'
  const quiz = isModel ? extractQuiz(text) : null

  if (quiz) {
    return { id: index, role: 'model', text: '', quiz }
  }

  return {
    id: index,
    role: isModel ? 'model' : 'user',
    text,
    file: raw.fileUri
      ? { name: 'Запись собеса', size: 0, status: 'done', uri: raw.fileUri }
      : undefined,
  }
}

export default function App() {
  const [chats, setChats] = useState([])
  const [threads, setThreads] = useState({})
  const [activeName, setActiveName] = useState(null)
  const [streamingId, setStreamingId] = useState(null)
  const [pendingDelete, setPendingDelete] = useState(null)
  const [creating, setCreating] = useState(false)
  const [error, setError] = useState(null)
  const [loading, setLoading] = useState(true)
  const [view, setView] = useState('vacancies')
  const [vacancies, setVacancies] = useState([])
  const [sources, setSources] = useState([])
  const [searching, setSearching] = useState(false)
  const [vacancyPage, setVacancyPage] = useState(1)
  const [vacancyTotal, setVacancyTotal] = useState(0)
  const [found, setFound] = useState(null)

  const abortRef = useRef(null)
  const nextId = useRef(1)

  const chat = chats.find((c) => c.name === activeName) ?? null
  const messages = threads[activeName] ?? []

  const handleError = useCallback((err) => {
    if (err.name === 'AbortError') return

    setError(
      err instanceof UnauthorizedError
        ? 'Не введён ключ Gemini — добавь GEMINI_API_KEY в .env'
        : 'Сервер недоступен: ' + err.message,
    )
  }, [])

  useEffect(() => {
    getChats()
      .then((list) => {
        const normalized = (list ?? []).map(normalizeChat)
        setChats(normalized)
        setActiveName((prev) => prev ?? normalized[0]?.name ?? null)
      })
      .catch(handleError)
      .finally(() => setLoading(false))
  }, [handleError])

  useEffect(() => {
    if (found) return

    getVacancies(PER_PAGE, (vacancyPage - 1) * PER_PAGE)
      .then((res) => {
        setVacancies(res?.vacancies ?? [])
        setVacancyTotal(res?.total ?? 0)
      })
      .catch(handleError)
  }, [vacancyPage, found, handleError])

  useEffect(() => {
    if (!activeName || threads[activeName]) return

    getMessages(activeName)
      .then((list) => {
        const normalized = (list ?? [])
          .filter((m) => !(m.text ?? '').includes(QUIZ_MARKER))
          .map(normalizeMessage)
        nextId.current = Math.max(nextId.current, normalized.length + 1)
        setThreads((prev) => ({ ...prev, [activeName]: normalized }))
      })
      .catch(handleError)
  }, [activeName, threads, handleError])

  function pushMessage(name, message) {
    setThreads((prev) => ({
      ...prev,
      [name]: [...(prev[name] ?? []), message],
    }))
  }

  function patchMessage(name, msgId, patch) {
    setThreads((prev) => ({
      ...prev,
      [name]: (prev[name] ?? []).map((m) =>
        m.id === msgId ? { ...m, ...patch } : m,
      ),
    }))
  }

  function touchChat(name, preview) {
    setChats((prev) =>
      prev.map((c) =>
        c.name === name ? { ...c, lastMessage: preview, updatedAt: 'сейчас' } : c,
      ),
    )
  }

  function stopStream() {
    abortRef.current?.abort()
    abortRef.current = null
    setStreamingId(null)
  }

  async function ask(name, req, replyId) {
    const controller = new AbortController()
    abortRef.current = controller
    setStreamingId(replyId)

    try {
      await streamMessage(name, req, {
        signal: controller.signal,
        onChunk: (chunk) =>
          setThreads((prev) => ({
            ...prev,
            [name]: prev[name].map((m) =>
              m.id === replyId ? { ...m, text: m.text + chunk } : m,
            ),
          })),
      })
    } catch (err) {
      handleError(err)
      patchMessage(name, replyId, { text: 'Не удалось получить ответ.' })
    } finally {
      abortRef.current = null
      setStreamingId(null)
    }
  }

  async function send(text, file) {
    const name = activeName
    if (!name) return

    const userId = nextId.current++
    const replyId = nextId.current++

    pushMessage(name, {
      id: userId,
      role: 'user',
      text,
      file: file ? { ...file, progress: 0 } : undefined,
    })
    touchChat(name, file ? 'Запись собеса' : text)

    let uploaded = null

    if (file) {
      const controller = new AbortController()
      abortRef.current = controller

      try {
        uploaded = await uploadInterview(name, file.raw, {
          signal: controller.signal,
          onProgress: (progress) =>
            patchMessage(name, userId, { file: { ...file, progress } }),
        })

        patchMessage(name, userId, {
          file: { ...file, status: 'in progress', uri: uploaded.uri },
        })
      } catch (err) {
        handleError(err)
        patchMessage(name, userId, { file: { ...file, status: 'failed' } })
        abortRef.current = null
        return
      }

      abortRef.current = null
    }

    pushMessage(name, { id: replyId, role: 'model', text: '' })

    await ask(
      name,
      {
        text: text || 'Разбери эту запись собеседования.',
        fileUri: uploaded?.uri,
        mimeType: uploaded?.mimeType,
      },
      replyId,
    )

    if (file) patchMessage(name, userId, { file: { ...file, status: 'done' } })
  }

  async function startQuiz() {
    const name = activeName
    if (!name) return

    const replyId = nextId.current++
    pushMessage(name, { id: replyId, role: 'model', text: '' })

    const controller = new AbortController()
    abortRef.current = controller
    setStreamingId(replyId)

    let raw = ''

    try {
      await streamMessage(
        name,
        { text: QUIZ_PROMPT },
        {
          signal: controller.signal,
          onChunk: (chunk) => {
            raw += chunk
          },
        },
      )
    } catch (err) {
      handleError(err)
      patchMessage(name, replyId, { text: 'Не удалось собрать вопросы.' })
      abortRef.current = null
      setStreamingId(null)
      return
    }

    abortRef.current = null
    setStreamingId(null)

    const questions = extractQuiz(raw)

    if (!questions) {
      patchMessage(name, replyId, {
        text: 'Не получилось разобрать вопросы, попробуй ещё раз.',
      })
      return
    }

    patchMessage(name, replyId, { text: 'Погоняю тебя по вакансии.' })
    pushMessage(name, { id: nextId.current++, role: 'model', text: '', quiz: questions })
    touchChat(name, 'Тренировка')
  }

  async function addChat(value) {
    setCreating(false)

    try {
      await createChat(value)
      setChats((prev) => [normalizeChat({ chatName: value }), ...prev])
      setThreads((prev) => ({ ...prev, [value]: [] }))
      setActiveName(value)
      setView('chat')
    } catch (err) {
      handleError(err)
    }
  }

  async function confirmDelete() {
    const name = pendingDelete.name
    stopStream()
    setPendingDelete(null)

    try {
      await deleteChat(name)
    } catch (err) {
      handleError(err)
      return
    }

    setChats((prev) => prev.filter((c) => c.name !== name))
    setThreads((prev) => {
      const next = { ...prev }
      delete next[name]
      return next
    })

    if (name === activeName) {
      setActiveName(chats.find((c) => c.name !== name)?.name ?? null)
    }
  }

  function resetVacancies() {
    setFound(null)
    setSources([])
    setVacancyPage(1)
  }

  function goVacancyPage(p) {
    const pageCount = Math.max(1, Math.ceil(vacancyTotal / PER_PAGE))
    const next = Math.min(Math.max(1, p), pageCount)

    setVacancyPage(next)

    if (found) {
      const from = (next - 1) * PER_PAGE
      setVacancies(found.slice(from, from + PER_PAGE))
    }
  }

  function pickChat(name) {
    stopStream()
    setActiveName(name)
    setView('chat')
  }

  async function runSearch(query) {
    setSearching(true)

    try {
      const result = await searchVacancies(query)
      setSources(result?.sources ?? [])

      const list = result?.vacancies ?? []
      setFound(list)
      setVacancyTotal(list.length)
      setVacancies(list.slice(0, PER_PAGE))
      setVacancyPage(1)
    } catch (err) {
      handleError(err)
    } finally {
      setSearching(false)
    }
  }

  async function chatAboutVacancy(vacancy) {
    const name = `${vacancy.company} — ${vacancy.title}`

    if (chats.some((c) => c.name === name)) {
      pickChat(name)
      return
    }

    try {
      await createChat(name)
    } catch (err) {
      handleError(err)
      return
    }

    setChats((prev) => [
      normalizeChat({
        chatName: name,
        company: vacancy.company,
        title: vacancy.title,
        updatedAt: 'сейчас',
      }),
      ...prev,
    ])
    setThreads((prev) => ({ ...prev, [name]: [] }))
    setActiveName(name)
    setView('chat')
  }

  return (
    <div className={s.app}>
      {error && (
        <div className={s.error} role="alert">
          {error}
          <button type="button" onClick={() => setError(null)}>
            ×
          </button>
        </div>
      )}

      <ChatList
        chats={chats}
        activeName={view === 'chat' ? activeName : null}
        loading={loading}
        showVacancies={view === 'vacancies'}
        onPick={pickChat}
        onCreate={() => setCreating(true)}
        onShowVacancies={() => setView('vacancies')}
      />

      {view === 'vacancies' ? (
        <VacancyBrowser
          vacancies={vacancies}
          sources={sources}
          searching={searching}
          page={vacancyPage}
          pageCount={Math.max(1, Math.ceil(vacancyTotal / PER_PAGE))}
          from={(vacancyPage - 1) * PER_PAGE}
          total={vacancyTotal}
          filtered={found !== null}
          onGo={goVacancyPage}
          onReset={resetVacancies}
          onSearch={runSearch}
          onCreateChat={chatAboutVacancy}
        />
      ) : (
        <ChatWindow
          chat={chat}
          messages={messages}
          streamingId={streamingId}
          onSend={send}
          onStop={stopStream}
          onDelete={() => setPendingDelete(chat)}
          onQuiz={startQuiz}
        />
      )}

      {creating && (
        <PromptDialog
          title="Новый чат"
          hint="Обычно это компания или вакансия — по этому названию чат хранит свой контекст."
          placeholder="Например: Ozon Tech"
          onConfirm={addChat}
          onCancel={() => setCreating(false)}
        />
      )}

      {pendingDelete && (
        <ConfirmDialog
          title="Удалить чат?"
          text={
            <>
              Переписка с <b>{pendingDelete.company}</b> будет удалена вместе со
              всем контекстом. Отменить не получится.
            </>
          }
          onConfirm={confirmDelete}
          onCancel={() => setPendingDelete(null)}
        />
      )}
    </div>
  )
}
