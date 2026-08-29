export const QUIZ_PROMPT = `Составь 5 вопросов для подготовки к этому собеседованию.
Верни ТОЛЬКО JSON без markdown-разметки и пояснений, строго в таком виде:
{"questions":[{"text":"вопрос","options":["вариант 1","вариант 2","вариант 3","вариант 4"],"correct":0,"explanation":"почему этот ответ верный"}]}
correct — индекс правильного варианта с нуля. Вопросы по стеку и задачам из контекста чата.`

export function extractQuiz(text) {
  const fenced = text.match(/```(?:json)?\s*([\s\S]*?)```/)
  const raw = fenced ? fenced[1] : text

  const start = raw.indexOf('{')
  const end = raw.lastIndexOf('}')
  if (start === -1 || end <= start) return null

  let parsed
  try {
    parsed = JSON.parse(raw.slice(start, end + 1))
  } catch {
    return null
  }

  const list = parsed?.questions
  if (!Array.isArray(list) || list.length === 0) return null

  const clean = list.filter(
    (q) =>
      typeof q?.text === 'string' &&
      Array.isArray(q.options) &&
      q.options.length >= 2 &&
      Number.isInteger(q.correct) &&
      q.correct >= 0 &&
      q.correct < q.options.length,
  )

  return clean.length > 0 ? clean : null
}
