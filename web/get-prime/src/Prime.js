import React, {useState, useEffect} from 'react'
import axios from 'axios'

function Prime() {
   
    const [num, setNum] = useState()
    const [getNumber, setGetNumber] = useState()
    const [checkNumber, setCheckNumber] = useState()
    const [prime, setPrime] = useState({})

    const handleGetPrimeClick = () => {
        setGetNumber(num)
        setNum()
    }

    const handleCheckPrimeClick = () => {
        setCheckNumber(num)
        setNum()
    }

    useEffect(() => {
        if (getNumber !== undefined) {   
            axios.get(`http://localhost:6060/v1/get_prime/${getNumber}`)
            .then(res => {
                console.log(res)
                setPrime(res.data)
            })
            .catch(err => {
                if (err.request.status === 400){
                    setPrime({error: true, msg: err.request.statusText})
                }else {
                    console.log(err)
                }
            })
        }
    }, [getNumber])

    useEffect(() => {   
        if (checkNumber !== undefined) {
            axios.get(`http://localhost:6060/v1/check_prime/${checkNumber}`)
            .then(res => {
                console.log(res)
                setPrime(res.data)
            })
            .catch(err => {
                if (err.request.status === 400){
                    setPrime({error: true, msg: err.request.statusText})
                }else {
                    console.log(err)
                }
            })
        }
    }, [checkNumber])

    return (
        <div>
            <input type = "text" value={num} onChange={e => setNum(e.target.value)}/>
            <br />
            <button type = "button" onClick={handleCheckPrimeClick}>Check If Prime</button>
            <button type = "button" onClick={handleGetPrimeClick}>Get Lower Prime</button>
            <br />
            <div>{prime.is_prime !== undefined && getNumber !== undefined && prime.is_prime && prime.prime_num && <span>Did you know? the number {getNumber} is also a prime number</span> }</div>
            <div>{prime.is_prime !== undefined && checkNumber !== undefined && prime.is_prime === false && <span>The number {checkNumber} is not a prime number</span> }</div>
            <div>{prime.is_prime !== undefined && checkNumber !== undefined && prime.is_prime === true && <span>The number {checkNumber} is a prime number</span> }</div>
            {prime.prime_num && getNumber !== undefined && <span>Highest prime number lower than {getNumber} is: {prime.prime_num}</span>}
            {prime.error && <span>Error for input - {getNumber}: {prime.msg}</span>}
        </div>
    )
}

export default Prime