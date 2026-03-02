import ReactMarkdown from "react-markdown"
import remarkBreaks from 'remark-breaks'

interface Props {
    value: string
    loading: boolean
}

/*
* Response parser
*/
function formatPrediction(text: string) {
    if (!text) return text

    const labels = ["Tempo médio:", "Pior caso:", "Justificativa:"]

    return text
        .split("\n")
        .map((line) => {
            const trimmed = line.trim()

            const match = labels.find(label => trimmed.startsWith(label))
            if (!match) return line

            const content = line.slice(match.length)
            return `**${match}**${content}`
        })
        .join("\n")
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
                    value === ""
                        ? <ReactMarkdown>No prediction generated yet.</ReactMarkdown>
                        : <ReactMarkdown remarkPlugins={[remarkBreaks]}>{formatPrediction(value)}</ReactMarkdown>
                )}
            </div>
        </div>
    )
}