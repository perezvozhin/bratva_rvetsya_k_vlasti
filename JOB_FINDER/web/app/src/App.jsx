import s from './App.module.css'

import Header from './components/Header/Header'
import Stats from './components/Stats/Stats'
import VacancyList from './components/VacancyList/VacancyList'
import Pipeline from './components/Pipeline/Pipeline'
import InterviewReview from './components/InterviewReview/InterviewReview'
import Events from './components/Events/Events'
import Helpers from './components/Helpers/Helpers'

import {
  MOCK_HEADER,
  MOCK_STATS,
  MOCK_VACANCIES,
  MOCK_PIPELINE,
  MOCK_INTERVIEW,
  MOCK_EVENTS,
  MOCK_HELPERS,
} from './mocks/data'

// данные пока из моков, запросы лежат в api.js
export default function App() {
  return (
    <div className={s.page}>
      <Header data={MOCK_HEADER} />
      <Stats items={MOCK_STATS} />

      <div className={s.layout}>
        <main className={s.column}>
          <VacancyList vacancies={MOCK_VACANCIES} />
          <Pipeline stages={MOCK_PIPELINE} />
        </main>

        <aside className={s.side}>
          <InterviewReview data={MOCK_INTERVIEW} />
          <Events items={MOCK_EVENTS} />
          <Helpers items={MOCK_HELPERS} />
        </aside>
      </div>
    </div>
  )
}
