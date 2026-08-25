import { RecordIcon } from '../../icons/Icons'
import s from './InterviewReview.module.css'
import common from '../../shared/common.module.css'

export default function InterviewReview({ data }) {
  return (
    <section className={common.block}>
      <h2 className={common.titleSm}>Разбор интервью</h2>

      <div className={s.panel}>
        <div className={s.head}>
          <span className={s.name}>{data.title}</span>
          <span className={s.meta}>{data.meta}</span>
        </div>

        <div className={s.insights}>
          {data.insights.map((item) => (
            <div key={item.id} className={s.insight}>
              <span className={`${s.score} ${s[item.tone]}`}>{item.score}</span>
              <span className={s.text}>{item.text}</span>
            </div>
          ))}
        </div>

        {/*нерабочая залупа на будущее*/}
        <button type="button" className={common.btnPrimary}>
          <RecordIcon />
          Записать следующий собес
        </button>
      </div>
    </section>
  )
}
