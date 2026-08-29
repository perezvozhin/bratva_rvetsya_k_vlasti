import { useEffect, useRef } from 'react'
import Message from '../Message/Message'
import Composer from '../Composer/Composer'
import Quiz from '../Quiz/Quiz'
import { TrashIcon, ChatIcon, QuizIcon } from '../../icons/Icons'
import s from './ChatWindow.module.css'
import common from '../../shared/common.module.css'

export default function ChatWindow({
  chat,
  messages,
  streamingId,
  onSend,
  onStop,
  onDelete,
  onQuiz,
}) {
  const scrollRef = useRef(null)

  useEffect(() => {
    const el = scrollRef.current
    if (el) el.scrollTop = el.scrollHeight
  }, [messages, streamingId])

  if (!chat) {
    return (
      <section className={s.window}>
        <div className={s.placeholder}>
          <span className={s.placeholderMark}>
            <ChatIcon size={40} />
          </span>
          <h2 className={s.placeholderTitle}>Выбери чат</h2>
          <p className={s.placeholderText}>
            Каждая компания живёт в своём чате со своим контекстом.
          </p>
        </div>
      </section>
    )
  }

  return (
    <section className={s.window}>
      <header className={s.head}>
        <div className={s.headBody}>
          <span className={s.company}>{chat.company}</span>
          <span className={s.title}>{chat.title}</span>
        </div>

        <div className={s.actions}>
          <button
            type="button"
            className={common.btnGhost}
            onClick={onQuiz}
            disabled={streamingId !== null}
          >
            <QuizIcon size={14} />
            Тренировка
          </button>

          <button
            type="button"
            className={common.iconBtn}
            onClick={onDelete}
            aria-label="Удалить чат"
          >
            <TrashIcon />
          </button>
        </div>
      </header>

      <div className={s.scroll} ref={scrollRef}>
        <div className={s.thread}>
          {messages.map((message) =>
            message.quiz ? (
              <Quiz key={message.id} questions={message.quiz} />
            ) : (
              <Message
                key={message.id}
                message={message}
                streaming={message.id === streamingId}
              />
            ),
          )}
        </div>
      </div>

      <Composer
        onSend={onSend}
        onStop={onStop}
        streaming={streamingId !== null}
      />
    </section>
  )
}
