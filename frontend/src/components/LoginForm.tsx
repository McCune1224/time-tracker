import React, { useState, useRef, useEffect } from 'react'

const LoginForm = () => {
    const userRef = useRef<string>("")
    const errRef = useRef()

    const [username, setUsername] = useState<string>()
    const [pwd, setPwd] = useState<string>()
    const [errMsg, setErrMsg] = useState<string>()
    const [success, setSuccess] = useState<boolean>(false)


    useEffect(() => {
        userRef.current.fixed

    }, [])

    return (
        <div></div>
    )
}

export default LoginForm
