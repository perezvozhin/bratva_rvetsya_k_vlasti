import { useEffect, useRef, useState } from 'react'
import s from './PromptDialog.module.css'
import common from '../../shared/common.module.css'

export default function PromptDialog({
  title,
  hint,
  placeholder,
  initialValue = '',
  confirmLabel = 'Создать',
  cancelLabel = 'Отмена',
  onConfirm,
  onCancel,
}) {
  const [value, setValue] = useState(initialValue)
  const inputRef = useRef(null)

  useEffect(() => {
    inputRef.current?.focus()

    function onKey(e) {
      if (e.key === 'Escape') onCancel()
    }

    document.addEventListener('keydown', onKey)
    return () => document.removeEventListener('keydown', onKey)
  }, [onCancel])

  function submit(e) {
    e.preventDefault()

    const trimmed = value.trim()
    if (!trimmed) return

    onConfirm(trimmed)
  }

  return (
    <div className={s.backdrop} onClick={onCancel} role="presentation">
      <form
        className={s.dialog}
        onClick={(e) => e.stopPropagation()}
        onSubmit={submit}
        aria-label={title}
      >
        <h2 className={s.title}>{title}</h2>
        {hint && <p className={s.hint}>{hint}</p>}

        <input
          ref={inputRef}
          className={s.input}
          value={value}
          onChange={(e) => setValue(e.target.value)}
          placeholder={placeholder}
        />

        <div className={s.actions}>
          <button type="button" className={common.btnGhost} onClick={onCancel}>
            {cancelLabel}
          </button>

          <button
            type="submit"
            className={common.btnPrimary}
            disabled={!value.trim()}
          >
            {confirmLabel}
          </button>
        </div>
      </form>
    </div>
  )
}
