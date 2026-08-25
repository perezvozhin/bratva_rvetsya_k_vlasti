// цвет и вердикт не хранятся в данных, оба выводятся из score
// границы 8 и 6 держать синхронно с бекендом

export function toneOf(score) {
  if (score >= 8) return 'green'
  if (score >= 6) return 'yellow'
  return 'red'
}

export function verdictOf(score) {
  if (score >= 8) return 'отличный вход'
  if (score >= 6) return 'надо попотеть'
  return 'много опыта'
}

// 9.4 -> "9,4"
export function formatScore(score) {
  return score.toFixed(1).replace('.', ',')
}
