import { PlusIcon, ChatIcon } from '../../icons/Icons'
import s from './ChatList.module.css'
import common from '../../shared/common.module.css'

export default function ChatList({ chats, activeId, onPick, onCreate }) {
  return (
    <aside className={s.sidebar}>
      <div className={s.head}>
        <div className={s.brand}>
          <span className={s.brandMark}>
            <ChatIcon size={17} />
          </span>
          Job Finder
        </div>

        <button
          type="button"
          className={`${common.btnGhost} ${s.newBtn}`}
          onClick={onCreate}
        >
          <PlusIcon size={14} />
          Новый чат
        </button>
      </div>

      <div className={s.scroll}>
        {chats.length === 0 ? (
          <p className={s.empty}>Чатов пока нет</p>
        ) : (
          chats.map((chat) => (
            <button
              type="button"
              key={chat.id}
              className={`${s.item} ${chat.id === activeId ? s.active : ''}`}
              onClick={() => onPick(chat.id)}
            >
              <span className={s.badge}>{chat.company.charAt(0)}</span>

              <span className={s.itemBody}>
                <span className={s.itemTop}>
                  <span className={s.company}>{chat.company}</span>
                  <span className={s.time}>{chat.updatedAt}</span>
                </span>
                <span className={s.title}>{chat.title}</span>
                <span className={s.preview}>{chat.lastMessage}</span>
              </span>
            </button>
          ))
        )}
      </div>
    </aside>
  )
}
