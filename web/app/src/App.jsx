import { useRef, useState } from 'react'
import ChatList from './components/ChatList/ChatList'
import ChatWindow from './components/ChatWindow/ChatWindow'
import ConfirmDialog from './components/ConfirmDialog/ConfirmDialog'
import { MOCK_CHATS, MOCK_MESSAGES } from './mocks/data'
import s from './App.module.css'

const FAKE_REPLY =
  'Пока это заглушка вместо Gemini. Здесь появится ответ модели с учётом контекста именно этого чата — компании, вакансии и всего, что вы уже обсудили.'

const FAKE_REVIEW =
  'Разобрал запись. Ответы по структуре держишь, но вступление затянуто — на первый вопрос ушло почти четыре минуты. По горутинам не хватило примеров из практики: теорию рассказал, а где сам ловил утечку, не показал. В конце не спросил про команду и процессы, а это обычно считают за незаинтересованность.'

export default function App() {
  const [chats, setChats] = useState(MOCK_CHATS)
  const [threads, setThreads] = useState(MOCK_MESSAGES)
  const [activeId, setActiveId] = useState(MOCK_CHATS[0]?.id ?? null)
  const [streamingId, setStreamingId] = useState(null)
  const [pendingDelete, setPendingDelete] = useState(null)

  const timerRef = useRef(null)
  const uploadRef = useRef(null)
  const nextId = useRef(1000)

  const chat = chats.find((c) => c.id === activeId) ?? null
  const messages = threads[activeId] ?? []

  function pushMessage(chatId, message) {
    setThreads((prev) => ({
      ...prev,
      [chatId]: [...(prev[chatId] ?? []), message],
    }))
  }

  function stopStream() {
    if (timerRef.current) clearInterval(timerRef.current)
    if (uploadRef.current) clearInterval(uploadRef.current)
    timerRef.current = null
    uploadRef.current = null
    setStreamingId(null)
  }

  function patchMessage(chatId, msgId, patch) {
    setThreads((prev) => ({
      ...prev,
      [chatId]: prev[chatId].map((m) =>
        m.id === msgId ? { ...m, ...patch } : m,
      ),
    }))
  }

  function streamReply(chatId, replyId, full) {
    setStreamingId(replyId)

    let i = 0
    timerRef.current = setInterval(() => {
      i += 2
      patchMessage(chatId, replyId, { text: full.slice(0, i) })
      if (i >= full.length) stopStream()
    }, 16)
  }

  function uploadFile(chatId, msgId, file) {
    let done = 0

    uploadRef.current = setInterval(() => {
      done += 7
      const progress = Math.min(done, 100)

      patchMessage(chatId, msgId, { file: { ...file, progress } })

      if (progress < 100) return

      clearInterval(uploadRef.current)
      uploadRef.current = null

      patchMessage(chatId, msgId, {
        file: { ...file, status: 'in progress' },
      })

      const replyId = nextId.current++
      pushMessage(chatId, { id: replyId, role: 'model', text: '' })

      setTimeout(() => {
        patchMessage(chatId, msgId, { file: { ...file, status: 'done' } })
        streamReply(chatId, replyId, FAKE_REVIEW)
      }, 700)
    }, 90)
  }

  function send(text, file) {
    const chatId = activeId
    const userId = nextId.current++

    pushMessage(chatId, {
      id: userId,
      role: 'user',
      text,
      file: file ? { ...file, progress: 0 } : undefined,
    })

    setChats((prev) =>
      prev.map((c) =>
        c.id === chatId
          ? {
              ...c,
              lastMessage: file ? 'Запись собеса' : text,
              updatedAt: 'сейчас',
            }
          : c,
      ),
    )

    if (file) {
      uploadFile(chatId, userId, file)
      return
    }

    const replyId = nextId.current++
    pushMessage(chatId, { id: replyId, role: 'model', text: '' })
    streamReply(chatId, replyId, FAKE_REPLY)
  }

  function createChat() {
    const id = nextId.current++
    const chat = {
      id,
      company: 'Новая компания',
      title: 'Без названия',
      lastMessage: 'Пустой чат',
      updatedAt: 'сейчас',
    }

    setChats((prev) => [chat, ...prev])
    setThreads((prev) => ({ ...prev, [id]: [] }))
    setActiveId(id)
  }

  function confirmDelete() {
    const id = pendingDelete.id
    stopStream()

    setChats((prev) => prev.filter((c) => c.id !== id))
    setThreads((prev) => {
      const next = { ...prev }
      delete next[id]
      return next
    })

    if (id === activeId) {
      const rest = chats.filter((c) => c.id !== id)
      setActiveId(rest[0]?.id ?? null)
    }

    setPendingDelete(null)
  }

  function pickChat(id) {
    stopStream()
    setActiveId(id)
  }

  return (
    <div className={s.app}>
      <ChatList
        chats={chats}
        activeId={activeId}
        onPick={pickChat}
        onCreate={createChat}
      />

      <ChatWindow
        chat={chat}
        messages={messages}
        streamingId={streamingId}
        onSend={send}
        onStop={stopStream}
        onDelete={() => setPendingDelete(chat)}
      />

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
