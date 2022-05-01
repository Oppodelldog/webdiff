const apiBaseUrl = "/rest"
const apiBaseUrlSessions = apiBaseUrl + "/sessions"
const apiBaseUrlSession = apiBaseUrl + "/session"
const apiBaseUrlFiles = apiBaseUrl + "/files"
const apiBaseUrlFile = apiBaseUrl + "/file"
const apiBaseUrlDiff = apiBaseUrl + "/diff"
const apiBaseUrlDownload = apiBaseUrl + "/download"

async function getSessions() {
    return await get(apiBaseUrlSessions)
}

async function getFiles(session) {
    return await get(apiBaseUrlFiles + "/" + session)
}

async function getFile(session, id) {
    return await get(apiBaseUrlFile + "/" + session + "/" + id)
}

async function getDiff(sessionA, idA, sessionB, idB) {
    return await get(apiBaseUrlDiff + `/${sessionA}/${idA}/${sessionB}/${idB}`)
}

async function getSessionUrls(session) {
    return await get(apiBaseUrlSession + `/${session}/urls`)
}

async function enqueueDownload(session, url) {
    await post(apiBaseUrlDownload, {
        url: url,
        session: session
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

    return response.json();
}

