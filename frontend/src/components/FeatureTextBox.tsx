import { useState } from "react"

interface FeatureTextBoxProps {
    value?: string
    onChange?: (value: string) => void
}

export default function FeatureTextBox({ value, onChange }: FeatureTextBoxProps) {
    const [internalValue, setInternalValue] = useState(value ?? "")

    const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
        const newValue = e.target.value
        setInternalValue(newValue)
        onChange?.(newValue)
    }

    return (
        <div>
            <p>Feature(s)</p>
            <textarea
                value={internalValue}
                onChange={handleChange}
                placeholder="Build a service to process payments asynchronously"
                rows={4}
            />
        </div>
    )
}
