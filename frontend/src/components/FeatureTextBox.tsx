import { useState } from "react"

interface FeatureTextBoxProps {
    value?: string
    onChange?: (value: string) => void
}

const MAX_LENGTH = 2000

export default function FeatureTextBox({ value, onChange }: FeatureTextBoxProps) {
    const [internalValue, setInternalValue] = useState(value ?? "")
    const charCount = internalValue.length
    const percentage = (charCount / MAX_LENGTH) * 100

    const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
        const newValue = e.target.value
        setInternalValue(newValue)
        onChange?.(newValue)
    }

    return (
        <div className="feature-textbox">
            <label>Feature(s):</label>

            <textarea
                value={internalValue}
                onChange={handleChange}
                placeholder="Build a service to process payments asynchronously"
                rows={4}
                maxLength={MAX_LENGTH}
            />

            <div className="char-counter">
                <span
                    className={`counter-text ${percentage > 90
                            ? "danger"
                            : percentage > 70
                                ? "warning"
                                : ""
                        }`}
                >
                    {charCount}/{MAX_LENGTH}
                </span>
            </div>
        </div>
    )
}