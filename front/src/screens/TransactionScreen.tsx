import { Button, Form } from "react-bootstrap"
import FormContainer from "../components/FormContainer"
import { SyntheticEvent, useState } from "react"

const TransactionScreen = () => {

    const [typeTrans, setTypeTrans] = useState('')
    const [nameProduct, setNameProduct] = useState('')
    const [quantity, setQuantity] = useState(0)
    const [sumMoney, setSumMoney] = useState(0)
    const [rate, setRate] = useState(0)

    const submitHandler = async(e: SyntheticEvent) => {
        e.preventDefault()
        
        const response = await fetch('http://localhost:8080/api/transaction/in', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: "include",
            body: JSON.stringify({
                type: typeTrans,
                name_product: nameProduct,
                quantity: quantity,
                sum_money: sumMoney
            })
        })

        const data = await response.json()
        setRate(data.rate)
        console.log(rate)
    }  

    return rate ? (
        <>
            <FormContainer>
                <Form.Group controlId="typeTrans" className="my-3">
                    <Form.Label>Тип транзакции</Form.Label>
                    <Form.Select aria-label="Default select example" value={typeTrans} 
                    onChange={t => setTypeTrans(t.target.value)}>
                        <option value="1">Купить</option>
                        <option value="2">Продать</option>
                    </Form.Select>
                </Form.Group>

                <Form.Group controlId="nameProduct" className="my-3">
                    <Form.Label>Название</Form.Label>
                    <Form.Control type='nameProduct' placeholder="Название" 
                    value={nameProduct}
                    onChange={n => setNameProduct(n.target.value)}
                    />
                </Form.Group>

                <Form.Group controlId="quantity" className="my-3">
                    <Form.Label>Количество</Form.Label>
                    <Form.Control type='quantity' placeholder="Количество" 
                    value={quantity}
                    onChange={q => setQuantity(Number(q.target.value))}
                    />
                </Form.Group>

                <Form.Group controlId="sumMoney" className="my-3">
                    <Form.Label>Сумма</Form.Label>
                    <Form.Control type='sumMoney' placeholder="Сумма" 
                    value={sumMoney}
                    onChange={s => setSumMoney(Number(s.target.value))}
                    />
                </Form.Group>

                <Form.Group controlId="rate" className="my-3">
                    <Form.Label>Курс</Form.Label>
                    <Form.Control type='sumMoney' value={rate}/>
                </Form.Group>

                <Button variant="primary" type="submit" onClick={submitHandler}>
                    Рассчитать
                </Button>
            </FormContainer>
        </>
    ) : (
        <>
            <FormContainer>
                <Form.Group controlId="typeTrans" className="my-3">
                    <Form.Label>Тип транзакции</Form.Label>
                    <Form.Select aria-label="Default select example" value={typeTrans}
                                 onChange={t => setTypeTrans(t.target.value)}>
                        <option value="1">Купить</option>
                        <option value="2">Продать</option>
                    </Form.Select>
                </Form.Group>

                <Form.Group controlId="nameProduct" className="my-3">
                    <Form.Label>Название</Form.Label>
                    <Form.Control type='nameProduct' placeholder="Название"
                                  value={nameProduct}
                                  onChange={n => setNameProduct(n.target.value)}
                    />
                </Form.Group>

                <Form.Group controlId="quantity" className="my-3">
                    <Form.Label>Количество</Form.Label>
                    <Form.Control type='quantity' placeholder="Количество"
                                  value={quantity}
                                  onChange={q => setQuantity(Number(q.target.value))}
                    />
                </Form.Group>

                <Form.Group controlId="sumMoney" className="my-3">
                    <Form.Label>Сумма</Form.Label>
                    <Form.Control type='sumMoney' placeholder="Сумма"
                                  value={sumMoney}
                                  onChange={s => setSumMoney(Number(s.target.value))}
                    />
                </Form.Group>

                <Button variant="primary" type="submit" onClick={submitHandler}>
                    Рассчитать
                </Button>
            </FormContainer>
        </>
    )
}

export default TransactionScreen