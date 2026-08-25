import s from './Header.module.css'
import common from '../../shared/common.module.css'

export default function Header({ data }) {
  return (
    <header className={s.header}>
      <div className={s.greeting}>
        <p className={s.date}>
          {data.today} · {data.freshCount} свежих вакансий
        </p>
        <h1 className={s.title}>
          Привет, {data.userName}. Есть на что посмотреть.
        </h1>
      </div>

      <div className={s.right}>
        {/* когда последний раз ходили на площадку */}
        <div className={s.sources}>
          {data.sources.map((item) => (
            <span key={item.id} className={s.source}>
              {item.label} · {item.lastRun}
            </span>
          ))}
        </div>

        <div className={common.avatar}>{data.userName.charAt(0)}</div>
      </div>
    </header>
  )
}
