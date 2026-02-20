import ReactMarkdown from "react-markdown"

interface Props {
    value: string
    loading: boolean
}

export default function PredictionField({ value, loading }: Props) {
    return (
        <div style={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
            <label>Prediction Output:</label>
            <div
                className="prediction-box"
                style={{
                    flex: 1,
                    overflow: "auto",
                    padding: "24px",
                    border: "1px solid var(--border-color)",
                    borderRadius: "12px",
                    background: "#09090b",
                    lineHeight: '1.6',
                    fontSize: '0.95rem'
                }}
            >
                {loading ? (
                    <div className="shimmer">Generating report...</div>
                ) : (
                    value === "" ? <ReactMarkdown>No prediction generated yet.</ReactMarkdown> : <ReactMarkdown>{value}</ReactMarkdown>
                )}
            </div>
        </div>
    )
}
