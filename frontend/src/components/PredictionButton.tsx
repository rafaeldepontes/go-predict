import styles from "./PredictionButton.module.css"

interface Props {
    loading: boolean,
    onClick: () => void,
}

export default function PredictionButton({ onClick, loading }: Props) {
    return (
        <button onClick={onClick} disabled={loading} className={styles.button}>
            {loading ? "Calculating..." : "Make Prediction"}
        </button>
    )
}