function formatDate(eventData) {
  const fmt = new Intl.DateTimeFormat('uk-UA', {
    dateStyle: 'full',
    timeStyle: 'short',
    timeZone: 'Europe/Kyiv',
  })
  if (eventData === undefined) return '...'
  if (eventData.is_onetime) {
    return fmt.format(new Date(eventData.date))
  } else {
    const time = new Date(eventData.time)
    const hour = time.getHours()
    const min = time.getMinutes()

    // Dude, this tongue is not ye olde C, why we do require these breaks? To break our faces

    let weekday = ''
    switch (eventData.weekday) {
      case 0:
        weekday = 'кожної неділі'
        break
      case 1:
        weekday = 'кожного понеділка'
        break
      case 2:
        weekday = 'кожного вівторка'
        break
      case 3:
        weekday = 'кожної середи'
        break
      case 4:
        weekday = 'кожного четверга'
        break
      case 5:
        weekday = "кожної п'ятниці"
        break
      case 6:
        weekday = 'кожної суботи'
        break
      default:
        weekday = 'кожного ніколи'
        break
    }

    return `${weekday} о ${String(hour).padStart(2, '0')}:${String(min).padStart(2, '0')}`
  }
}

function formatVisitDate(visit) {
  const fmt = new Intl.DateTimeFormat('uk-UA', {
    dateStyle: 'full',
    timeStyle: 'short',
    timeZone: 'Europe/Kyiv',
  })
  return fmt.format(new Date(visit.visit_date))
}

export { formatDate, formatVisitDate }
