import styles from "./ErrorPopUp.module.css"

interface ErrorPopUpProps {
    error: string,
    onClose: () => void
}

export default function ErrorPopUp({ error, onClose }: ErrorPopUpProps) {
    return (
        <div className={styles["error-popup"]}>
            <div className={styles["error-content"]}>
                <span>{error}</span>
                <button className={styles["error-close"]} onClick={onClose}>
                    x
                </button>
            </div>
        </div>
    )
}
