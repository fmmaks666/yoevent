import { normalizeData } from '../utils/utils.js'

const API = import.meta.env.VITE_BACKEND_URL

async function getEvents(visitor = undefined) {
  return fetch(`${API}/events?visitor=${visitor}`)
}

async function getRegistrations(visitor = undefined) {
  return fetch(`${API}/registrations?visitor=${visitor}`)
}

async function getEvent(eventId = undefined) {
  return fetch(`${API}/event?event=${eventId}`)
}

async function getVisits(visitor = undefined) {
  return fetch(`${API}/visits?visitor=${visitor}`)
}

async function postVisit(eventId, visitor) {
  return fetch(`${API}/visit`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ event_id: eventId, visitor: visitor }),
  })
}

async function postEventRegistration(registrationId, visitor) {
  return fetch(`${API}/registration`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ registration_id: registrationId, visitor: visitor }),
  })
}

async function deleteEventRegistration(registrationId, visitor) {
  const params = new URLSearchParams({
    registration_id: registrationId,
    visitor: visitor,
  })

  return fetch(`${API}/registration?${params.toString()}`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
  })
}

async function postVisitor(data) {
  const req = normalizeData(data)
  const json = JSON.stringify(req)
  return fetch(`${API}/visitor`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: json,
  })
}

async function updateVisitor(data) {
  const req = normalizeData(data)
  const json = JSON.stringify(req)
  return fetch(`${API}/visitor`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: json,
  })
}

async function deleteVisitor(data) {
  data.sex = Number(data.sex)
  const json = JSON.stringify(data)
  return fetch(`${API}/visitor`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: json,
  })
}

async function checkAuth(token) {
  return fetch(`${API}/admin/check`, {
    headers: {
      Authorization: `Basic ${token}`,
    },
  })
}

async function getEventsAdmin(token) {
  return fetch(`${API}/admin/events`, {
    headers: {
      Authorization: `Basic ${token}`,
    },
  })
}

async function getRegistrationsAdmin(token) {
  return fetch(`${API}/admin/registrations`, {
    headers: {
      Authorization: `Basic ${token}`,
    },
  })
}

async function getVisitors(token) {
  return fetch(`${API}/admin/visitors`, {
    headers: {
      Authorization: `Basic ${token}`,
    },
  })
}

async function postEvent(token, data) {
  const json = JSON.stringify(data)
  return fetch(`${API}/admin/event`, {
    method: 'POST',
    headers: {
      Authorization: `Basic ${token}`,
    },
    body: json,
  })
}

async function updateEvent(token, data) {
  const json = JSON.stringify(data)
  return fetch(`${API}/admin/event`, {
    method: 'PUT',
    headers: {
      Authorization: `Basic ${token}`,
    },
    body: json,
  })
}

async function doDeleteEvent(token, eventId) {
  const params = new URLSearchParams({
    event_id: eventId,
  })
  return fetch(`${API}/admin/event?${params.toString()}`, {
    method: 'DELETE',
    headers: {
      Authorization: `Basic ${token}`,
    },
  })
}

async function downloadVisits(token, eventId, date, all = false) {
  const params = new URLSearchParams({
    event_id: eventId,
    all: all,
    month: date.month,
    year: date.year,
  })
  return fetch(`${API}/admin/visits?${params.toString()}`, {
    headers: {
      Authorization: `Basic ${token}`,
    },
  })
}

async function getEventStats(token, eventId, date, all = false) {
  const params = new URLSearchParams({
    event_id: eventId,
    all: all,
    month: date.month,
    year: date.year,
  })
  return fetch(`${API}/admin/stats?${params.toString()}`, {
    headers: {
      Authorization: `Basic ${token}`,
    },
  })
}

export {
  getEvents,
  getEvent,
  getVisits,
  postVisit,
  postVisitor,
  updateVisitor,
  checkAuth,
  getEventsAdmin,
  getVisitors,
  postEvent,
  updateEvent,
  downloadVisits,
  getEventStats,
  doDeleteEvent,
  getRegistrations,
  postEventRegistration,
  deleteEventRegistration,
  getRegistrationsAdmin,
}
