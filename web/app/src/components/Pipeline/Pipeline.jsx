import s from './Pipeline.module.css'
import common from '../../shared/common.module.css'

export default function Pipeline({ stages }) {
  return (
    <section className={common.block}>
      <h2 className={common.title}>Отклики</h2>

      <div className={s.pipeline}>
        {stages.map((stage) => (
          <div key={stage.id} className={s.stage}>
            <span className={s.count}>{stage.count}</span>
            <span className={s.label}>{stage.label}</span>
            <div className={s.bar}>
              {/* ширина приходит с бека */}
              <div
                className={`${s.fill} ${s[stage.tone]}`}
                style={{ width: stage.width }}
              />
            </div>
          </div>
        ))}
      </div>
    </section>
  )
}
