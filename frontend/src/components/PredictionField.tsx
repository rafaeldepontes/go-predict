import ReactMarkdown from "react-markdown"

interface Props {
    value: string
    loading: boolean
}

export default function PredictionField({ value, loading }: Props) {
    const content = loading ? "Thinking..." : value;

    return (
        <div style={{ flex: 1 }}>
            <p>Prediction:</p>

            <div
                style={{
                    width: "100%",
                    height: "70%",
                    overflow: "auto",
                    padding: "12px",
                    border: "1px solid #ccc",
                    borderRadius: "6px",
                    background: "#181818",
                    textAlign: "left"
                }}
            >
                <ReactMarkdown>{content}</ReactMarkdown>
            </div>
        </div>
    )
}
