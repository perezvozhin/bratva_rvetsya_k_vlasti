// запросы к своему go-серверу
// пути относительные: в проде фронт внутри бинарника, в деве vite проксирует /api

/*
export async function fetchJSON(path) {
  const resp = await fetch(path)

  // 401 = ключ гемини не введен, надо показать экран ввода
  // смотри логику в httpmw.ApiCheckMiddleWare
  if (resp.status === 401) return null

  if (!resp.ok) throw new Error('bad status: ' + resp.status)
  return resp.json()
}

export const getVacancies = () => fetchJSON('/api/vacancies')
export const getStats = () => fetchJSON('/api/stats')
export const getPipeline = () => fetchJSON('/api/pipeline')
export const getInterview = () => fetchJSON('/api/interviews')
export const getEvents = () => fetchJSON('/api/events')
export const getHelpers = () => fetchJSON('/api/helpers')

// отметка должна пережить перезагрузку, поэтому пишем в базу
export async function saveMark(id, marked) {
  await fetch('/api/vacancies/' + id + '/mark', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ marked }),
  })
}
*/
