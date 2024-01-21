// src\utils\storage.ts
export const getItem = <T>(key: string) => {
    const data = window.localStorage.getItem(key)
    if (!data) return null
    try {
        return JSON.parse(data) as T
    } catch (error) {
        return null
    }
}

export const setItem = (key: string, value: Object | string | null) => {
    if (typeof value !== 'string') {
        value = JSON.stringify(value)
    }
    window.localStorage.setItem(key, value as string)
}

export const removeItem = (key: string) => {
    window.localStorage.removeItem(key)
}

