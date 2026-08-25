import { CheckIcon, PlusIcon } from '../../icons/Icons.jsx'
import { toneOf, verdictOf, formatScore } from '../../shared/score.js'
import s from './VacancyCard.module.css'

export default function VacancyCard({ vacancy, marked, onToggle }) {
  // модули дают хешированные имена, поэтому класс по ключу
  const tone = s[toneOf(vacancy.score)]

  return (
    <article className={s.card}>
      <span className={`${s.indicator} ${tone}`} />

      <div className={s.body}>
        <div className={s.head}>
          <span className={s.title}>{vacancy.title}</span>
          <span className={s.company}>{vacancy.company}</span>
        </div>

        <div className={s.meta}>
          <span className={s.salary}>{vacancy.salary}</span>
          <span>опыт {vacancy.exp}</span>
          <span>{vacancy.city}</span>
          <span>{vacancy.tagline}</span>
        </div>
      </div>

      <div className={s.side}>
        <div className={s.score}>
          <span className={s.scoreValue}>{formatScore(vacancy.score)}</span>
          <span className={s.verdict}>{verdictOf(vacancy.score)}</span>
        </div>

        <button
          type="button"
          className={`${s.mark} ${marked ? s.markOn : ''}`}
          onClick={onToggle}
          aria-pressed={marked}
          aria-label={marked ? 'Убрать из интересных' : 'В интересные'}
        >
          {marked ? <CheckIcon /> : <PlusIcon />}
        </button>
      </div>
    </article>
  )
}
