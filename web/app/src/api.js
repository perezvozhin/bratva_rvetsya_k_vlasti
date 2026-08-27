/*
export async function fetchJSON(path, opts) {
  const resp = await fetch(path, opts)

  if (resp.status === 401) return null

  if (!resp.ok) throw new Error('bad status: ' + resp.status)
  return resp.json()
}

export const getChats = () => fetchJSON('/api/chats')
export const getMessages = (chatId) => fetchJSON('/api/chats/' + chatId + '/messages')

export const createChat = (company, title) =>
  fetchJSON('/api/chats', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ company, title }),
  })

export const deleteChat = (chatId) =>
  fetchJSON('/api/chats/' + chatId, { method: 'DELETE' })

export function uploadInterview(chatId, file, { onProgress, signal }) {
  return new Promise((resolve, reject) => {
    const form = new FormData()
    form.append('file', file)

    const xhr = new XMLHttpRequest()
    xhr.open('POST', '/api/chats/' + chatId + '/interview')

    xhr.upload.onprogress = (e) => {
      if (e.lengthComputable) onProgress(Math.round((e.loaded / e.total) * 100))
    }

    xhr.onload = () =>
      xhr.status < 400
        ? resolve(JSON.parse(xhr.responseText))
        : reject(new Error('bad status: ' + xhr.status))

    xhr.onerror = () => reject(new Error('upload failed'))
    signal?.addEventListener('abort', () => xhr.abort())

    xhr.send(form)
  })
}

export async function streamMessage(chatId, text, { onChunk, signal }) {
  const resp = await fetch('/api/chats/' + chatId + '/messages', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ text }),
    signal,
  })

  if (!resp.ok) throw new Error('bad status: ' + resp.status)

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

      onChunk(JSON.parse(payload).text)
    }
  }
}
*/
