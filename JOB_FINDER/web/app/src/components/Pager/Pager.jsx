import { ArrowLeftIcon, ArrowRightIcon } from '../../icons/Icons'
import s from './Pager.module.css'

export default function Pager({ from, shownCount, total, page, pageCount, onGo }) {
  return (
    <div className={s.pager}>
      <span className={s.range}>
        {shownCount > 0
          ? `${from + 1}–${from + shownCount} из ${total}`
          : `0 из ${total}`}
      </span>

      <div className={s.controls}>
        <button
          type="button"
          className={s.btn}
          onClick={() => onGo(page - 1)}
          disabled={page === 1}
          aria-label="Предыдущая страница"
        >
          <ArrowLeftIcon />
        </button>

        {Array.from({ length: pageCount }, (_, i) => i + 1).map((p) => (
          <button
            type="button"
            key={p}
            className={`${s.btn} ${p === page ? s.on : ''}`}
            onClick={() => onGo(p)}
          >
            {p}
          </button>
        ))}

        <button
          type="button"
          className={s.btn}
          onClick={() => onGo(page + 1)}
          disabled={page === pageCount}
          aria-label="Следующая страница"
        >
          <ArrowRightIcon />
        </button>
      </div>
    </div>
  )
}
