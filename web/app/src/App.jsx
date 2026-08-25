import s from './App.module.css'

import Header from './components/Header/Header.jsx'
import Stats from './components/Stats/Stats.jsx'
import VacancyList from './components/VacancyList/VacancyList.jsx'
import Pipeline from './components/Pipeline/Pipeline.jsx'
import InterviewReview from './components/InterviewReview/InterviewReview.jsx'
import Events from './components/Events/Events.jsx'
import Helpers from './components/Helpers/Helpers.jsx'

import {
  MOCK_HEADER,
  MOCK_STATS,
  MOCK_VACANCIES,
  MOCK_PIPELINE,
  MOCK_INTERVIEW,
  MOCK_EVENTS,
  MOCK_HELPERS,
} from './mocks/data.js'

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
