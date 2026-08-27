import { useEffect, useRef } from 'react'
import s from './ConfirmDialog.module.css'
import common from '../../shared/common.module.css'

export default function ConfirmDialog({
  title,
  text,
  confirmLabel = 'Удалить',
  cancelLabel = 'Отмена',
  onConfirm,
  onCancel,
}) {
  const confirmRef = useRef(null)

  useEffect(() => {
    confirmRef.current?.focus()

    function onKey(e) {
      if (e.key === 'Escape') onCancel()
    }

    document.addEventListener('keydown', onKey)
    return () => document.removeEventListener('keydown', onKey)
  }, [onCancel])

  return (
    <div
      className={s.backdrop}
      onClick={onCancel}
      role="presentation"
    >
      <div
        className={s.dialog}
        role="alertdialog"
        aria-modal="true"
        aria-label={title}
        onClick={(e) => e.stopPropagation()}
      >
        <h2 className={s.title}>{title}</h2>
        <div className={s.text}>{text}</div>

        <div className={s.actions}>
          <button type="button" className={common.btnGhost} onClick={onCancel}>
            {cancelLabel}
          </button>

          <button
            type="button"
            ref={confirmRef}
            className={s.danger}
            onClick={onConfirm}
          >
            {confirmLabel}
          </button>
        </div>
      </div>
    </div>
  )
}
