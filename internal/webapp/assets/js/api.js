const apiBaseUrl = "/rest"
const apiBaseUrlSessions = apiBaseUrl + "/session"
const apiBaseUrlFiles = apiBaseUrl + "/files"
const apiBaseUrlDiff = apiBaseUrl + "/diff"

async function getSessions() {
    return await get(apiBaseUrlSessions)
}

async function getFiles(session) {
    return await get(apiBaseUrlFiles + "/" + session)
}

async function getDiff(sessionA, idA, sessionB, idB) {
    return await get(apiBaseUrlDiff + `/${sessionA}/${idA}/${sessionB}/${idB}`)
}

async function get(url) {
    const res = await fetch(url);
    if (!res.ok) {
        const message = `An error has occurred: ${res.status} - ${res.statusText}`;
        throw new Error(message);
    }
    return await res.json();
}
