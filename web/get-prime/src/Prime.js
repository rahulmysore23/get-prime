import React, {useState, useEffect} from 'react'
import axios from 'axios'

function Prime() {
   
    const [num, setNum] = useState(2)
    const [getNumber, setGetNumber] = useState(2)
    const [checkNumber, setCheckNumber] = useState(2)
    const [prime, setPrime] = useState({})
    // const [displayPrime, setDisplayPrime] = useState(true)

    const handleGetPrimeClick = () => {
        setGetNumber(num)
    }

    const handleCheckPrimeClick = () => {
        setCheckNumber(num)
    }

    useEffect(() => {   
        axios.get(`http://localhost:6060/v1/get_prime/${getNumber}`)
        .then(res => {
            console.log(res)
            setPrime(res.data)
        })
        .catch(err => {
            console.log(err)
        })
    }, [getNumber])

    useEffect(() => {   
        axios.get(`http://localhost:6060/v1/check_prime/${checkNumber}`)
        .then(res => {
            console.log(res)
            setPrime(res.data)
        })
        .catch(err => {
            console.log(err)
        })
    }, [checkNumber])

    return (
        <div>
            <input type = "text" value={num} onChange={e => setNum(e.target.value)}/>
            <br />
            <button type = "button" onClick={handleCheckPrimeClick}>Check If Prime</button>
            <button type = "button" onClick={handleGetPrimeClick}>Get Lower Prime</button>
            <br />
            <div>{'' + prime.is_prime}</div>
            <div>{prime.prime_num && <h1>Highest prime number lower than {num} is: {prime.prime_num}</h1>}</div>
        </div>
    )
}

export default Prime