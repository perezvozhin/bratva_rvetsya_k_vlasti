import { useState } from 'react'
import Pager from './Pager'
import { ChatIcon, CheckIcon } from '../../icons/Icons'
import s from './VacancyBrowser.module.css'
import common from '../../shared/common.module.css'

export default function VacancyBrowser({
  vacancies,
  sources,
  searching,
  page,
  pageCount,
  from,
  total,
  filtered,
  onGo,
  onSearch,
  onReset,
  onCreateChat,
}) {
  const [text, setText] = useState('golang')
  const [onlyRemote, setOnlyRemote] = useState(false)

  function submit(e) {
    e.preventDefault()
    onSearch({ text: text.trim(), onlyRemote, page: 1, perPage: 20 })
  }

  return (
    <section className={s.browser}>
      <div className={s.head}>
        <h2 className={s.title}>Поиск вакансий</h2>

        <form className={s.form} onSubmit={submit}>
          <input
            className={s.search}
            value={text}
            onChange={(e) => setText(e.target.value)}
            placeholder="Например: golang backend"
          />

          <label className={onlyRemote ? s.toggleOn : s.toggle}>
            <input
              type="checkbox"
              className={s.hidden}
              checked={onlyRemote}
              onChange={(e) => setOnlyRemote(e.target.checked)}
            />
            <span className={onlyRemote ? s.boxOn : s.box}>
              <CheckIcon size={11} />
            </span>
            только удалённо
          </label>

          <button
            type="submit"
            className={common.btnPrimary}
            disabled={searching || !text.trim()}
          >
            {searching ? 'Ищу…' : 'Искать'}
          </button>
        </form>

        {sources.length > 0 && (
          <div className={s.sources}>
            {sources.map((src) => (
              <span
                key={src.source}
                className={src.error ? s.sourceBad : undefined}
                title={src.error}
              >
                {src.source} · {src.error ? 'недоступен' : src.found + ' найдено'}
              </span>
            ))}

            {filtered && (
              <button type="button" className={s.reset} onClick={onReset}>
                показать всё, что собрано
              </button>
            )}
          </div>
        )}
      </div>

      <div className={s.scroll}>
        {vacancies.length === 0 ? (
          <p className={s.empty}>
            {searching ? 'Ищу вакансии…' : 'Пока пусто — запусти поиск.'}
          </p>
        ) : (
          <div className={s.list}>
            {vacancies.map((v) => (
              <article key={v.id} className={s.card}>
                <div className={s.body}>
                  <div className={s.top}>
                    <a
                      className={s.name}
                      href={v.url}
                      target="_blank"
                      rel="noopener noreferrer"
                    >
                      {v.title}
                    </a>
                    <span className={s.company}>{v.company}</span>
                  </div>

                  <div className={s.meta}>
                    <span className={s.salary}>{v.salary}</span>
                    {v.experience && <span>{v.experience}</span>}
                    {v.remote && <span>удалённо</span>}
                    <span className={s.tag}>{v.source}</span>
                  </div>
                </div>

                <div className={s.side}>
                  <button
                    type="button"
                    className={common.btnGhost}
                    onClick={() => onCreateChat(v)}
                  >
                    <ChatIcon size={14} />
                    Обсудить
                  </button>
                </div>
              </article>
            ))}
          </div>
        )}

        <Pager
          page={page}
          pageCount={pageCount}
          from={from}
          shown={vacancies.length}
          total={total}
          onGo={onGo}
        />
      </div>
    </section>
  )
}
