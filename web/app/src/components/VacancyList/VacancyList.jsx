import { useState } from 'react'
import VacancyCard from '../VacancyCard/VacancyCard.jsx'
import Pager from '../Pager/Pager.jsx'
import s from './VacancyList.module.css'
import common from '../../shared/common.module.css'
// import { saveMark } from '../../api'

const PER_PAGE = 5
const FILTERS = ['Все', 'Топ', 'Интересные', 'LinkedIn']

export default function VacancyList({ vacancies }) {
  const [filter, setFilter] = useState('Все')
  const [marked, setMarked] = useState({ 2: true })
  const [page, setPage] = useState(1)

  function toggleMark(id) {
    setMarked((prev) => ({ ...prev, [id]: !prev[id] }))
    // saveMark(id, !marked[id])
  }

  function pickFilter(label) {
    setFilter(label)
    setPage(1)
  }

  let list = vacancies
  if (filter === 'Топ') list = list.filter((v) => v.score >= 8)
  if (filter === 'Интересные') list = list.filter((v) => marked[v.id])
  if (filter === 'LinkedIn') list = list.filter((v) => v.source === 'linkedin')

  const sorted = [...list].sort((a, b) => b.score - a.score)

  const pageCount = Math.max(1, Math.ceil(sorted.length / PER_PAGE))
  const current = Math.min(page, pageCount)
  const from = (current - 1) * PER_PAGE
  const shown = sorted.slice(from, from + PER_PAGE)

  const go = (p) => setPage(Math.min(Math.max(1, p), pageCount))

  return (
    <section className={common.block}>
      <div className={common.head}>
        <h2 className={common.title}>Вакансии</h2>

        <nav className={s.filters}>
          {FILTERS.map((label) => (
            <button
              type="button"
              key={label}
              className={`${s.filter} ${label === filter ? s.on : ''}`}
              onClick={() => pickFilter(label)}
            >
              {label}
            </button>
          ))}
        </nav>
      </div>

      {/* фиксированная высота, чтобы пагинация не прыгала на неполной странице */}
      <div className={s.body}>
        {shown.length === 0 ? (
          <p className={s.empty}>Ничего не нашлось.</p>
        ) : (
          <div className={s.list}>
            {shown.map((vacancy) => (
              <VacancyCard
                key={vacancy.id}
                vacancy={vacancy}
                marked={!!marked[vacancy.id]}
                onToggle={() => toggleMark(vacancy.id)}
              />
            ))}
          </div>
        )}
      </div>

      <Pager
        from={from}
        shownCount={shown.length}
        total={sorted.length}
        page={current}
        pageCount={pageCount}
        onGo={go}
      />
    </section>
  )
}
