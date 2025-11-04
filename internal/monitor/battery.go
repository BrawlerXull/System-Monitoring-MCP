package monitor

import "github.com/distatus/battery"

// GetBatteryStatus returns information about all system batteries
func GetBatteryStatus() ([]map[string]interface{}, error) {
	batteries, err := battery.GetAll()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for _, b := range batteries {
		result = append(result, map[string]interface{}{
			"design_capacity": b.Design,
			"full_capacity":   b.Full,
			"current":         b.Current,
			"charge_percent":  (b.Current / b.Full) * 100,
			"state":           b.State.String(),
			"voltage":         b.Voltage,
		})
	}
	return result, nil
}
