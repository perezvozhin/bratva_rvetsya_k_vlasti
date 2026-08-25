import s from './Stats.module.css'

export default function Stats({ items }) {
  return (
    <section className={s.stats}>
      {items.map((item) => (
        <div key={item.id} className={s.stat}>
          <span className={s.label}>{item.label}</span>
          <span className={s.value}>{item.value}</span>
        </div>
      ))}
    </section>
  )
}
