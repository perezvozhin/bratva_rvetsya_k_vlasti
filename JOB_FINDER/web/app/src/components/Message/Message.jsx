import Attachment from '../Attachment/Attachment'
import s from './Message.module.css'

export default function Message({ message, streaming }) {
  const isUser = message.role === 'user'

  return (
    <div className={isUser ? s.user : s.row}>
      <span className={isUser ? s.avatarUser : s.avatarModel}>
        {isUser ? 'Я' : 'G'}
      </span>

      <div className={isUser ? s.stackUser : s.stack}>
        {message.file && (
          <Attachment file={message.file} inBubble={isUser} />
        )}

        {(message.text || streaming) && (
          <div className={isUser ? s.bubbleUser : s.bubbleModel}>
            {message.text}
            {streaming && <span className={s.caret} />}
          </div>
        )}
      </div>
    </div>
  )
}
