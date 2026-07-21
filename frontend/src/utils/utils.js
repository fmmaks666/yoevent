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
    // THE WERID ASS TRICKS TO GET THIS BS WORKING
    // TODO: KILL MYSELF OR MAKE THIS IN ANY GOOD WAY RESILIENT
    const hour = time.getHours() + 3
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
  const damn = new Date(visit.visit_date)
  damn.setHours(damn.getHours() + 3)
  return fmt.format(damn)
}

// TO THE PERSON READING THIS
// I LOVE YOU BECAUSE I HATE THIS
// THE BACKEND, MY DEAR BACKEND SENDS THE OFFSET
// AND YET IT IGNORES IT COMPLETELY GOING FOR F'KING +00:00
// WHY IS THAT IS NOT CLEAR TO ME BUT I WANT TO KILL EVERYTHING OR RATHER GO CRY PLAYING CLANNAD
// DON'T YOU THINK IT'S FUNNY HOW A JAPANESE BISHOUJO GAME CAN MAKE A UKRAINIAN GUY CRY?
// LOOK THIS WORKS ON MY PC WHERE I HAVE CLANNAD INSTALLED
// IT MAKES THE DATE CORRECT. THOUGH I WONDER WHAT WILL HAPPEN WHEN I DEPLOY TO MY SERVER IN
// GERMANY, WHICH DOESN'T HAVE CLANNAD INSTALLED
function formatBirthdate(date) {
  const fmt = new Intl.DateTimeFormat('uk-UA', {
    timeZone: 'Europe/Kyiv',
    day: 'numeric',
    month: 'long',
    year: 'numeric',
  })
  console.log(date)
  console.log(new Date(date).toLocaleString())
  const fuck = new Date(date)
  fuck.setHours(fuck.getHours() + 3)
  return fmt.format(fuck)
}

export { formatDate, formatVisitDate, formatBirthdate }
