import { Button, Modal } from "react-bootstrap"

interface Props {
    show: boolean
    uid: any
    handleClose(): void
}

const ModalWindow = ({ show, uid, handleClose }: Props) => {
    return(
        <Modal show={show} onHide={handleClose}>
            <Modal.Header closeButton>
            <Modal.Title>Идентификационный id</Modal.Title>
            </Modal.Header>
            <Modal.Body>Скопируйте: {uid}</Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={handleClose}>
                    Скопировать
                </Button>
            </Modal.Footer>
        </Modal>
    )
}

export default ModalWindow