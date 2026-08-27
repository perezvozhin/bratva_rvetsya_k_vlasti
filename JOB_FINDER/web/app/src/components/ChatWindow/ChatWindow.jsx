import { useEffect, useRef } from 'react'
import Message from '../Message/Message'
import Composer from '../Composer/Composer'
import { TrashIcon, ChatIcon } from '../../icons/Icons'
import s from './ChatWindow.module.css'
import common from '../../shared/common.module.css'

export default function ChatWindow({
  chat,
  messages,
  streamingId,
  onSend,
  onStop,
  onDelete,
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

        <button
          type="button"
          className={common.iconBtn}
          onClick={() => onDelete(chat.id)}
          aria-label="Удалить чат"
        >
          <TrashIcon />
        </button>
      </header>

      <div className={s.scroll} ref={scrollRef}>
        <div className={s.thread}>
          {messages.map((message) => (
            <Message
              key={message.id}
              message={message}
              streaming={message.id === streamingId}
            />
          ))}
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
