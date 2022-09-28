import axios from 'axios'
import React, { useState } from 'react'

const RandomNumber = () => {
    const [randNumber, setRandomNumber] = useState<number>()

    const handleClick = () => {
        axios.get("http://localhost:8080/rng")
            .then((response) => setRandomNumber(response.data.NUMBER))
    }

    return (
        <div>
            <button 
                className="font-bold bg-red-300 hover:bg-red-500 active:bg-red-700 text-white"
                onClick={handleClick}>Random Number</button>
            <h1>{randNumber}</h1>
        </div>
    )
}

export default RandomNumber
