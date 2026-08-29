import { useRef, useState } from 'react'
import { SendIcon, StopIcon, ClipIcon } from '../../icons/Icons'
import Attachment from '../Attachment/Attachment'
import s from './Composer.module.css'
import common from '../../shared/common.module.css'

const ACCEPT = 'video/mp4,video/quicktime,video/webm'

export default function Composer({ onSend, onStop, streaming, disabled }) {
  const [text, setText] = useState('')
  const [file, setFile] = useState(null)
  const [dragging, setDragging] = useState(false)

  const inputRef = useRef(null)
  const fileRef = useRef(null)

  function resize() {
    const el = inputRef.current
    if (!el) return
    el.style.height = 'auto'
    el.style.height = el.scrollHeight + 'px'
  }

  function change(e) {
    setText(e.target.value)
    resize()
  }

  function pickFile(picked) {
    if (!picked) return
    setFile({ name: picked.name, size: picked.size, status: 'new', raw: picked })
  }

  function drop(e) {
    e.preventDefault()
    setDragging(false)
    pickFile(e.dataTransfer.files[0])
  }

  function submit() {
    const value = text.trim()
    if ((!value && !file) || streaming || disabled) return

    onSend(value, file)
    setText('')
    setFile(null)
    if (fileRef.current) fileRef.current.value = ''

    const el = inputRef.current
    if (el) el.style.height = 'auto'
  }

  function keyDown(e) {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault()
      submit()
    }
  }

  const empty = !text.trim() && !file

  return (
    <div className={s.composer}>
      <div
        className={dragging ? s.wrapDrag : s.wrap}
        onDragOver={(e) => {
          e.preventDefault()
          setDragging(true)
        }}
        onDragLeave={() => setDragging(false)}
        onDrop={drop}
      >
        {file && (
          <div className={s.file}>
            <Attachment file={file} />
            <button
              type="button"
              className={s.remove}
              onClick={() => setFile(null)}
              aria-label="Убрать запись"
            >
              ×
            </button>
          </div>
        )}

        <div className={s.line}>
          <input
            ref={fileRef}
            type="file"
            accept={ACCEPT}
            className={s.hidden}
            onChange={(e) => pickFile(e.target.files[0])}
          />

          <button
            type="button"
            className={s.clip}
            onClick={() => fileRef.current?.click()}
            disabled={disabled || streaming}
            aria-label="Прикрепить запись собеса"
          >
            <ClipIcon />
          </button>

          <textarea
            ref={inputRef}
            className={s.input}
            rows={1}
            value={text}
            onChange={change}
            onKeyDown={keyDown}
            disabled={disabled}
            placeholder="Спросить или закинуть запись собеса…"
          />

          {streaming ? (
            <button
              type="button"
              className={`${common.btnPrimary} ${s.send}`}
              onClick={onStop}
              aria-label="Остановить"
            >
              <StopIcon />
            </button>
          ) : (
            <button
              type="button"
              className={`${common.btnPrimary} ${s.send}`}
              onClick={submit}
              disabled={empty || disabled}
              aria-label="Отправить"
            >
              <SendIcon />
            </button>
          )}
        </div>
      </div>

      <p className={s.hint}>
        Enter — отправить, Shift + Enter — перенос строки
      </p>
    </div>
  )
}
