interface Props {
    value: string
    loading: boolean
}

export default function PredictionField({ value, loading }: Props) {
    return (
        <div style={{ flex: 1 }}>
            <p>Prediction:</p>
            <textarea
                value={loading ? "Thinking..." : value}
                readOnly
                style={{ width: "100%", height: "70%" }}
            />
        </div>
    )
}
