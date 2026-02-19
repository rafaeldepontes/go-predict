interface ErrorPopUpProps {
    error: string,
    onClose: () => void
}

export default function ErrorPopUp({ error, onClose }: ErrorPopUpProps) {
    return (
        <div className="error-popup">
            <div className="error-content">
                <span>{error}</span>
                <button className="error-close" onClick={onClose}>
                    x
                </button>
            </div>
        </div>
    )
}
