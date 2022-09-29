import axios from 'axios'
import React, { useState, useEffect } from 'react'

const RandomNumber = () => {
    const [randNumber, setRandomNumber] = useState<number>()
    const [isPending, setIsPending] = useState<boolean>(true)


    useEffect(() => {
        setTimeout(() => {
            axios.get("http://localhost:8080/rng")
                .then((response) => {
                    setRandomNumber(response.data.number)
                }
                )
                .then(() => setIsPending(false))

        }, 1000)
    }, [])

    return (
        <div>
            {isPending && <div>Loading...</div>}
            <h1>{randNumber}</h1>
        </div>
    )
}

export default RandomNumber
