export class UnauthorizedError extends Error {}

async function request(path, opts) {
  const resp = await fetch(path, opts)

  if (resp.status === 401) throw new UnauthorizedError('no api key')
  if (!resp.ok) throw new Error('bad status: ' + resp.status)

  return resp
}

async function fetchJSON(path, opts) {
  const resp = await request(path, opts)
  const text = await resp.text()
  return text ? JSON.parse(text) : null
}

export const getChats = () => fetchJSON('/api/chats')

export const getMessages = (chatName) =>
  fetchJSON('/api/chats/' + encodeURIComponent(chatName) + '/messages')

export const createChat = (name) =>
  fetchJSON('/api/chats', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ chatName: name }),
  })

export const deleteChat = (chatName) =>
  fetchJSON('/api/chats/' + encodeURIComponent(chatName), { method: 'DELETE' })

export function uploadInterview(chatName, file, { onProgress, signal } = {}) {
  return new Promise((resolve, reject) => {
    const form = new FormData()
    form.append('file', file)

    const xhr = new XMLHttpRequest()
    xhr.open(
      'POST',
      '/api/chats/' + encodeURIComponent(chatName) + '/interview',
    )

    xhr.upload.onprogress = (e) => {
      if (e.lengthComputable && onProgress) {
        onProgress(Math.round((e.loaded / e.total) * 100))
      }
    }

    xhr.onload = () => {
      if (xhr.status === 401) return reject(new UnauthorizedError('no api key'))
      if (xhr.status >= 400) return reject(new Error('bad status: ' + xhr.status))

      try {
        resolve(JSON.parse(xhr.responseText))
      } catch (err) {
        reject(err)
      }
    }

    xhr.onerror = () => reject(new Error('upload failed'))
    xhr.onabort = () => reject(new DOMException('aborted', 'AbortError'))

    signal?.addEventListener('abort', () => xhr.abort())

    xhr.send(form)
  })
}

export async function streamMessage(chatName, req, { onChunk, signal } = {}) {
  const resp = await request(
    '/api/chats/' + encodeURIComponent(chatName) + '/messages',
    {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        text: req.text,
        fileUri: req.fileUri ?? '',
        MIMEType: req.mimeType ?? '',
      }),
      signal,
    },
  )

  const reader = resp.body.pipeThrough(new TextDecoderStream()).getReader()
  let buffer = ''

  while (true) {
    const { done, value } = await reader.read()
    if (done) break

    buffer += value
    const parts = buffer.split('\n\n')
    buffer = parts.pop()

    for (const part of parts) {
      const line = part.split('\n').find((l) => l.startsWith('data: '))
      if (!line) continue

      const payload = line.slice(6)
      if (payload === '[DONE]') return

      onChunk(JSON.parse(payload).text ?? '')
    }
  }
}

export const getVacancies = (limit = 50, offset = 0, q = '') =>
  fetchJSON(
    '/api/vacancies?limit=' + limit + '&offset=' + offset +
      (q ? '&q=' + encodeURIComponent(q) : ''),
  )

export const searchVacancies = (query) =>
  fetchJSON('/api/vacancies/search', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(query),
  })
