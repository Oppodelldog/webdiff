const apiBaseUrl = "/rest"
const apiBaseUrlSessions = apiBaseUrl + "/sessions"
const apiBaseUrlSession = apiBaseUrl + "/session"
const apiBaseUrlFiles = apiBaseUrl + "/files"
const apiBaseUrlFile = apiBaseUrl + "/file"
const apiBaseUrlFilter = apiBaseUrl + "/filter"
const apiBaseUrlFilters = apiBaseUrl + "/filters"
const apiBaseUrlDiff = apiBaseUrl + "/diff"
const apiBaseUrlDownload = apiBaseUrl + "/download"

async function getSessions() {
    return await get(apiBaseUrlSessions)
}

async function getFiles(session) {
    return await get(apiBaseUrlFiles + "/" + session)
}

async function getFile(session, id, filterName, prettify) {
    return await get(apiBaseUrlFile + "/" + session + "/" + id + query(prettify, filterName))
}

async function getFilters() {
    return await get(apiBaseUrlFilters)
}

async function upsertFilter(name, def, type) {
    return await post(apiBaseUrlFilter, {name, def, type})
}

async function deleteFilter(name) {
    return await del(apiBaseUrlFilter + "/" + name)
}

async function getDiff(sessionA, idA, sessionB, idB, filterName, prettify) {
    return await get(apiBaseUrlDiff + `/${sessionA}/${idA}/${sessionB}/${idB}${query(prettify, filterName)}`)
}

async function getSessionUrls(session) {
    return await get(apiBaseUrlSession + `/${session}/urls`)
}

async function enqueueDownload(session, url, fileId) {
    await post(apiBaseUrlDownload, {
        url: url,
        session: session,
        id: fileId
    })
}

async function get(url) {
    const res = await fetch(url);
    if (!res.ok) {
        const message = `An error has occurred: ${res.status} - ${res.statusText}`;
        throw new Error(message);
    }
    return await res.json();
}

async function post(url, data) {
    const response = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    });

    if (response.status === 204) {
        return
    }

    return response.json();
}

async function del(url) {
    await fetch(url, {method: 'DELETE',});
}


function query(prettify, filterName) {
    let query = "?pretty=" + prettify;
    if (filterName != null && filterName.length > 0) {
        query += `&filter=${filterName}`
    }

    return query
}