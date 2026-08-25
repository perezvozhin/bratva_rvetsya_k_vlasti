// GET /api/vacancies
export const MOCK_VACANCIES = [
  { id: 1, title: 'Go Developer', company: 'Yandex Cloud', salary: '220–280 тыс ₽', exp: 'не нужен', city: 'Москва, гибрид', tagline: 'Go · PostgreSQL', score: 9.4, source: 'hh' },
  { id: 2, title: 'Backend Engineer, Go', company: 'Ozon Tech', salary: '250–300 тыс ₽', exp: '1 год', city: 'удалённо', tagline: 'Go · Kafka', score: 8.1, source: 'hh' },
  { id: 3, title: 'Junior Golang Dev', company: 'Tinkoff', salary: '180–210 тыс ₽', exp: '1–3 года', city: 'Москва', tagline: 'Go · Redis', score: 6.7, source: 'hh' },
  { id: 4, title: 'Software Engineer', company: 'Revolut', salary: '€45–58k', exp: '2 года', city: 'Remote EU', tagline: 'Go · AWS · англ B2', score: 6.2, source: 'linkedin' },
  { id: 5, title: 'Middle Go Developer', company: 'Sber', salary: '260 тыс ₽', exp: '3–6 лет', city: 'Москва', tagline: 'Go · Oracle', score: 4.3, source: 'hh' },
  { id: 6, title: 'Go Developer (стажировка+)', company: 'Avito', salary: '190–240 тыс ₽', exp: 'не нужен', city: 'удалённо', tagline: 'Go · gRPC', score: 8.9, source: 'hh' },
]

export const MOCK_HEADER = {
  userName: 'Игорь',
  today: 'Понедельник, 25 августа',
  freshCount: 248,
  sources: [{ id: 'hh', label: 'hh.ru', lastRun: '14 мин назад' }],
}

// GET /api/stats
export const MOCK_STATS = [
  { id: 'inWork', label: 'Откликов в работе', value: '34' },
]

// GET /api/pipeline 
export const MOCK_PIPELINE = [
  { id: 'sent', label: 'отправлено', count: 34, width: '100%', tone: 'neutral' },
]

// GET /api/interviews
export const MOCK_INTERVIEW = {
  title: 'Ozon · Go Middle',
  meta: '42 мин · вчера',
  insights: [
    { id: 1, score: '7/10', tone: 'good', text: 'Ответы по STAR держишь, но вступление затянуто.' },
  ],
}

// GET /api/events
export const MOCK_EVENTS = [
  { id: 1, date: '26 авг', title: 'Ozon · тех. собес', meta: '15:00 · Go, конкурентность' },
]

// GET /api/helpers
export const MOCK_HELPERS = [
  { id: 1, name: 'Дима К.', drops: 14, thanks: 63, karma: 840 },
]
