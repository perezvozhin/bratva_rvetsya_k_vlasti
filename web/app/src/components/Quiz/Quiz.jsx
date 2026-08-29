import { useEffect, useState } from 'react'
import { playCorrect, playWrong } from '../../shared/sound'
import s from './Quiz.module.css'
import common from '../../shared/common.module.css'

const KEYS = ['A', 'B', 'C', 'D', 'E']

export default function Quiz({ questions, onFinish }) {
  const [index, setIndex] = useState(0)
  const [picked, setPicked] = useState(null)
  const [right, setRight] = useState(0)

  const question = questions[index]
  const answered = picked !== null
  const isRight = answered && picked === question.correct

  useEffect(() => {
    if (!answered) return
    if (isRight) playCorrect()
    else playWrong()
  }, [answered, isRight])

  function pick(i) {
    if (answered) return

    setPicked(i)
    if (i === question.correct) setRight((n) => n + 1)
  }

  function next() {
    if (index + 1 >= questions.length) {
      onFinish?.({ right, total: questions.length })
      return
    }

    setIndex((i) => i + 1)
    setPicked(null)
  }

  function optionClass(i) {
    if (!answered) return s.option
    if (i === question.correct) return `${s.option} ${s.correct}`
    if (i === picked) return `${s.option} ${s.wrong}`
    return s.option
  }

  return (
    <div className={s.quiz}>
      <div className={s.head}>
        <span className={s.label}>Тренировка</span>
        <span className={s.progress}>
          {index + 1} / {questions.length}
        </span>
      </div>

      <p className={s.question}>{question.text}</p>

      <div className={s.options}>
        {question.options.map((option, i) => (
          <button
            type="button"
            key={option}
            className={optionClass(i)}
            onClick={() => pick(i)}
            disabled={answered}
          >
            <span className={s.key}>{KEYS[i]}</span>
            {option}
          </button>
        ))}
      </div>

      {answered && (
        <div className={isRight ? s.feedbackGood : s.feedbackBad}>
          <span
            className={`${s.verdict} ${isRight ? s.verdictGood : s.verdictBad}`}
          >
            {isRight ? 'Верно' : 'Неверно'}
          </span>
          {question.explanation}
        </div>
      )}

      {answered && (
        <div className={s.actions}>
          <button type="button" className={common.btnPrimary} onClick={next}>
            {index + 1 >= questions.length ? 'Завершить' : 'Дальше'}
          </button>
        </div>
      )}
    </div>
  )
}
