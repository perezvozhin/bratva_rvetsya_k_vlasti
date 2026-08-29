import { ArrowLeftIcon, ArrowRightIcon } from '../../icons/Icons'
import s from './VacancyBrowser.module.css'

function pageList(page, total) {
  if (total <= 7) return Array.from({ length: total }, (_, i) => i + 1)

  const out = [1]
  const from = Math.max(2, page - 1)
  const to = Math.min(total - 1, page + 1)

  if (from > 2) out.push('…')
  for (let p = from; p <= to; p++) out.push(p)
  if (to < total - 1) out.push('…')

  out.push(total)
  return out
}

export default function Pager({ page, pageCount, from, shown, total, onGo }) {
  if (total === 0) return null

  return (
    <div className={s.pager}>
      <span className={s.range}>
        {from + 1}–{from + shown} из {total}
      </span>

      <div className={s.pagerControls}>
        <button
          type="button"
          className={s.pageBtn}
          onClick={() => onGo(page - 1)}
          disabled={page === 1}
          aria-label="Предыдущая страница"
        >
          <ArrowLeftIcon />
        </button>

        {pageList(page, pageCount).map((p, i) =>
          p === '…' ? (
            <span key={'dots' + i} className={s.dots}>
              …
            </span>
          ) : (
            <button
              type="button"
              key={p}
              className={p === page ? s.pageOn : s.pageBtn}
              onClick={() => onGo(p)}
            >
              {p}
            </button>
          ),
        )}

        <button
          type="button"
          className={s.pageBtn}
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
