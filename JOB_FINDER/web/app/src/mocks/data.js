// GET /api/chats
export const MOCK_CHATS = [
  {
    id: 1,
    company: 'Ozon Tech',
    title: 'Backend Engineer, Go',
    lastMessage: 'Спроси про их подход к очередям',
    updatedAt: '14:20',
  },
  {
    id: 2,
    company: 'Yandex Cloud',
    title: 'Go Developer',
    lastMessage: 'Разобрали вопросы по конкурентности',
    updatedAt: 'вчера',
  },
  {
    id: 3,
    company: 'Tinkoff',
    title: 'Junior Golang Dev',
    lastMessage: 'Что подтянуть перед собесом',
    updatedAt: '23 авг',
  },
]

// GET /api/chats/:id/messages
export const MOCK_MESSAGES = {
  1: [
    { id: 1, role: 'user', text: 'Что спросить у Ozon на техническом собесе?' },
    {
      id: 2,
      role: 'model',
      text: 'Спроси, как у них устроена работа с Kafka: свои обёртки или чистый клиент, как решают проблему повторной обработки. Это покажет, что ты думаешь про надёжность, а не только про код.',
    },
    { id: 3, role: 'user', text: 'А про кодстайл стоит?' },
    {
      id: 4,
      role: 'model',
      text: 'Стоит, но не про табы и пробелы. Спроси про ревью: сколько апрувов нужно, есть ли линтер в CI, кто мёржит. Это важнее.',
    },
  ],
  2: [
    {
      id: 1,
      role: 'user',
      text: 'Глянь, как я отвечал',
      file: {
        name: 'yandex-screening.mp4',
        size: 184 * 1024 * 1024,
        duration: '38 мин',
        status: 'done',
      },
    },
    {
      id: 2,
      role: 'model',
      text: 'Посмотрел запись. Держишься уверенно, но на вопросе про каналы ушёл в теорию вместо примера. В следующий раз начинай с случая из практики, потом обобщай.',
    },
    { id: 3, role: 'user', text: 'Какие вопросы по горутинам обычно задают?' },
    {
      id: 4,
      role: 'model',
      text: 'Чаще всего: чем канал отличается от мьютекса, что такое утечка горутины и как её поймать, зачем нужен context. Последнее спрашивают почти всегда.',
    },
  ],
  3: [{ id: 1, role: 'user', text: 'С чего начать подготовку?' }],
}
