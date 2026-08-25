import s from './Helpers.module.css'
import common from '../../shared/common.module.css'

export default function Helpers({ items }) {
  return (
    <section className={common.block}>
      <h2 className={common.titleSm}>Спасибо на этой неделе</h2>

      <div className={s.helpers}>
        {items.map((item) => (
          <div key={item.id} className={s.helper}>
            {/* аватарки нет, берем первую букву имени */}
            <div className={common.avatarSm}>{item.name.charAt(0)}</div>
            <div className={s.body}>
              <span className={s.name}>{item.name}</span>
              <span className={s.meta}>
                {item.drops} вакансий · {item.thanks} спасибо
              </span>
            </div>
            <span className={s.karma}>{item.karma}</span>
          </div>
        ))}
      </div>

      <p className={s.note}>
        Кинул хорошую вакансию через бота — джуны говорят спасибо, карма растёт.
      </p>
    </section>
  )
}
