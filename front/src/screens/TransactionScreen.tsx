import { Button, Form } from "react-bootstrap"
import FormContainer from "../components/FormContainer"
import { useState } from "react"

const TransactionScreen = () => {

    const [typeTrans, setTypeTrans] = useState('')
    const [nameProduct, setNameProduct] = useState('')
    const [quantity, setQuantity] = useState<any>()
    const [sumMoney, setSumMoney] = useState<any>()

    return (
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
                    onChange={q => setQuantity(q.target.value)}
                    />
                </Form.Group>

                <Form.Group controlId="sumMoney" className="my-3">
                    <Form.Label>Сумма</Form.Label>
                    <Form.Control type='sumMoney' placeholder="Сумма" 
                    value={sumMoney}
                    onChange={s => setSumMoney(s.target.value)}
                    />
                </Form.Group>

                <Button variant="primary" type="submit" onClick={undefined}>
                    Рассчитать
                </Button>
            </FormContainer>
        </>
    )
}

export default TransactionScreen