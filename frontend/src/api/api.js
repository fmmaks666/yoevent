import { normalizeData } from '../utils/utils.js'

const API = import.meta.env.VITE_BACKEND_URL

async function getEvents(visitor = undefined) {
  return fetch(`${API}/events?visitor=${visitor}`)
}

async function getEvent(eventId = undefined) {
  return fetch(`${API}/event?event=${eventId}`)
}

async function getVisits(visitor = undefined) {
  return fetch(`${API}/visits?visitor=${visitor}`)
}

async function postVisit(event_id, visitor) {
  return fetch(`${API}/visit`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ event_id: event_id, visitor: visitor }),
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
}
