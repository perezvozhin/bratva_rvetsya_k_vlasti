import { VideoIcon } from '../../icons/Icons'
import s from './Attachment.module.css'

const STATUS_LABEL = {
  new: 'загружено',
  'in progress': 'анализирую…',
  done: 'разобрано',
}

const STATUS_CLASS = {
  new: s.new,
  'in progress': s.progress,
  done: s.done,
}

function formatSize(bytes) {
  const mb = bytes / (1024 * 1024)
  if (mb >= 1024) return (mb / 1024).toFixed(1).replace('.', ',') + ' ГБ'
  return mb.toFixed(1).replace('.', ',') + ' МБ'
}

export default function Attachment({ file, inBubble }) {
  const uploading = file.progress !== undefined && file.progress < 100

  return (
    <div className={inBubble ? s.inBubble : s.card}>
      <span className={s.mark}>
        <VideoIcon size={17} />
      </span>

      <div className={s.body}>
        <span className={s.name}>{file.name}</span>

        <span className={s.meta}>
          {formatSize(file.size)}
          {file.duration ? ` · ${file.duration}` : ''}
          {' · '}
          <span className={`${s.status} ${STATUS_CLASS[file.status] ?? ''}`}>
            {uploading
              ? `${file.progress}%`
              : (STATUS_LABEL[file.status] ?? file.status)}
          </span>
        </span>

        {uploading && (
          <div className={s.bar}>
            <div className={s.fill} style={{ width: file.progress + '%' }} />
          </div>
        )}
      </div>
    </div>
  )
}
