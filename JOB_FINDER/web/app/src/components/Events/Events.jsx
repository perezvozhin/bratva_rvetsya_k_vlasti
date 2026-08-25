import s from './Events.module.css'
import common from '../../shared/common.module.css'

export default function Events({ items }) {
  return (
    <section className={common.block}>
      <h2 className={common.titleSm}>Скоро</h2>

      <div className={s.events}>
        {items.map((item) => (
          <div key={item.id} className={s.event}>
            <span className={s.date}>{item.date}</span>
            <div className={s.body}>
              <span className={s.title}>{item.title}</span>
              <span className={s.meta}>{item.meta}</span>
            </div>
            <button type="button" className={common.btnGhost}>
              подготовиться
            </button>
          </div>
        ))}
      </div>
    </section>
  )
}
