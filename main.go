package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Endpoint to run the planets binary
	app.Get("/run-planets", func(c *fiber.Ctx) error {
		birthdate := c.Query("birthdate")
		utctime := c.Query("utctime")
		latitude := c.Query("latitude")
		longitude := c.Query("longitude")
		altitude := c.Query("altitude")
		housesystem := c.Query("housesystem")

		formattedDate, err := parseBirthdate(birthdate)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(map[string]string{
				"error": fmt.Sprintf("Failed to parse birthdate: %v", err),
			})
		}

		args := fmt.Sprintf("-b%s -utc%s -p0123456789m -fPlbs -hsy%s -geopos%s,%s,%s -g\",\" -head", formattedDate, utctime, housesystem, longitude, latitude, altitude)

		output, err := runBinary("swetest", args)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(map[string]string{
				"error": fmt.Sprintf("Failed to run planets binary: %v", err),
			})
		}

		parsedOutput, err := parsePlanetsOutput(output)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(map[string]string{
				"error": fmt.Sprintf("Failed to parse planets binary output: %v", err),
			})
		}

		return c.JSON(parsedOutput)
	})

	// Endpoint to run the houses binary
	app.Get("/run-houses", func(c *fiber.Ctx) error {
		birthdate := c.Query("birthdate")
		utctime := c.Query("utctime")
		latitude := c.Query("latitude")
		longitude := c.Query("longitude")
		altitude := c.Query("altitude")
		housesystem := c.Query("housesystem")

		formattedDate, err := parseBirthdate(birthdate)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(map[string]string{
				"error": fmt.Sprintf("Failed to parse birthdate: %v", err),
			})
		}

		args := fmt.Sprintf("-house -p -fPlb -b%s -utc%s -hsy%s -geopos%s,%s,%s -g\",\" -head", formattedDate, utctime, housesystem, longitude, latitude, altitude)

		output, err := runBinary("swetest", args)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(map[string]string{
				"error": fmt.Sprintf("Failed to run houses binary: %v", err),
			})
		}

		parsedOutput, err := parseHousesOutput(output)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(map[string]string{
				"error": fmt.Sprintf("Failed to parse houses binary output: %v", err),
			})
		}

		return c.JSON(parsedOutput)
	})

// Endpoint to run the star binary
app.Get("/run-star", func(c *fiber.Ctx) error {
	birthdate := c.Query("birthdate")
	utctime := c.Query("utctime")
	stars := c.Query("stars")

	formattedDate, err := parseBirthdate(birthdate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": fmt.Sprintf("Failed to parse birthdate: %v", err),
		})
	}

	starList := strings.Split(stars, ",")

	var result []map[string]string

	for _, star := range starList {
		args := fmt.Sprintf("-b%s -utc%s -pf -fPlbsjw= -xf%s -head -g\",\"", formattedDate, utctime, star)

		output, err := runBinary("swetest", args)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(map[string]string{
				"error": fmt.Sprintf("Failed to run star binary for star '%s': %v", star, err),
			})
		}

		parsedOutput, err := parseStarOutput(output)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(map[string]string{
				"error": fmt.Sprintf("Failed to parse star binary output for star '%s': %v", star, err),
			})
		}

		result = append(result, parsedOutput...)
	}
	fmt.Println("Star result:", result)
	return c.JSON(result)
})


	// Endpoint to get the available options and parameters
	app.Get("/options", func(c *fiber.Ctx) error {
		response := getOptionResponse()
		return c.JSON(response)
	})

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app.Listen(":" + port)
}

func parseBirthdate(date string) (string, error) {
	formats := []string{"2.1.2006", "2-1-2006", "2/1/2006"}
	var parsedTime time.Time
	var err error

	for _, format := range formats {
		parsedTime, err = time.Parse(format, date)
		if err == nil {
			break
		}
	}

	if err != nil {
		return "", fmt.Errorf("Failed to parse birthdate: %v", err)
	}

	// Validate the month value
	if parsedTime.Month() > 12 {
		return "", fmt.Errorf("Invalid month value in birthdate: %d", parsedTime.Month())
	}

	formattedDate := parsedTime.Format("2.1.2006")

	return formattedDate, nil
}

func runBinary(binaryName string, args string) (string, error) {
	cmd := exec.Command(binaryName, strings.Split(args, " ")...)
    fmt.Println("This commmand:", cmd)

	outputPipe, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("Failed to create output pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("Failed to start command: %v", err)
	}

	outputBytes, err := ioutil.ReadAll(outputPipe)
	if err != nil {
		return "", fmt.Errorf("Failed to read command output: %v", err)
	}

	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("Command failed: %v", err)
	}

	output := string(outputBytes)

	fmt.Println("Binary Output:", output)
	// fmt.Println("Some error:", err)

	return output, nil
}

func parsePlanetsOutput(output string) ([]map[string]string, error) {
	lines := strings.Split(output, "\n")
	var result []map[string]string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || !strings.Contains(line, ",") {
			continue
		}

		fields := strings.Split(line, ",")
		if len(fields) != 4 {
			return nil, fmt.Errorf("Failed to parse planets binary output line: %s", line)
		}

		name := strings.TrimSpace(strings.Trim(fields[0], "\""))           // Remove leading/trailing spaces and extra `"`
		longitude := strings.TrimSpace(strings.Trim(fields[1], "\""))      // Remove leading/trailing spaces and extra `"`
		latitude := strings.TrimSpace(strings.Trim(fields[2], "\""))       // Remove leading/trailing spaces and extra `"`
		dailySpeed := strings.TrimSpace(strings.Trim(fields[3], "\""))     // Remove leading/trailing spaces and extra `"`

		item := map[string]string{
			"name":       name,
			"longitude":  longitude,
			"latitude":   latitude,
			"dailySpeed": dailySpeed,
		}

		result = append(result, item)
	}

	return result, nil
}

	
	func parseHousesOutput(output string) ([]map[string]string, error) {
		lines := strings.Split(output, "\n")
		var result []map[string]string
	
		for _, line := range lines {
			line = strings.TrimSpace(line)
	
			if line == "" || !strings.Contains(line, ",") {
				continue
			}
	
			fields := strings.Split(line, ",")
			if len(fields) != 2 {
				return nil, fmt.Errorf("Failed to parse houses binary output line: %s", line)
			}
	
			name := strings.TrimSpace(strings.Trim(fields[0], "\""))      // Remove leading/trailing spaces and extra `"`
			longitude := strings.TrimSpace(strings.Trim(fields[1], "\"")) // Remove leading/trailing spaces and extra `"`
	
			item := map[string]string{
				"name":      name,
				"longitude": longitude,
			}
	
			result = append(result, item)
		}
	
		return result, nil
	}
	
	
	
	func parseStarOutput(output string) ([]map[string]string, error) {
		lines := strings.Split(output, "\n")
		var result []map[string]string
	
		for _, line := range lines {
			line = strings.TrimSpace(line)
	
			if line == "" || !strings.Contains(line, ",") {
				continue
			}
	
			fields := strings.Split(line, ",")
			if len(fields) != 8 {
				return nil, fmt.Errorf("Failed to parse star output line: %s", line)
			}
	
			starName := strings.TrimSpace(strings.Trim(fields[0], "\""))  // Remove leading/trailing spaces and extra `"`
			altName := strings.TrimSpace(strings.Trim(fields[1], "\""))   // Remove leading/trailing spaces and extra `"`
			longitude := strings.TrimSpace(strings.Trim(fields[2], "\"")) // Remove leading/trailing spaces and extra `"`
			latitude := strings.TrimSpace(strings.Trim(fields[3], "\""))  // Remove leading/trailing spaces and extra `"`
			speed := strings.TrimSpace(strings.Trim(fields[4], "\""))     // Remove leading/trailing spaces and extra `"`
			house := strings.TrimSpace(strings.Trim(fields[5], "\""))     // Remove leading/trailing spaces and extra `"`
			distance := strings.TrimSpace(strings.Trim(fields[6], "\""))  // Remove leading/trailing spaces and extra `"`
			magnitude := strings.TrimSpace(strings.Trim(fields[7], "\"")) // Remove leading/trailing spaces and extra `"`
			magnitude = strings.TrimSuffix(magnitude, "m")                // Strip trailing "m"
	
			item := map[string]string{
				"starName":  starName,
				"altName":   altName,
				"longitude": longitude,
				"latitude":  latitude,
				"speed":     speed,
				"house":     house,
				"distance":  distance,
				"magnitude": magnitude,
			}
	
			result = append(result, item)
		}
	
		return result, nil
	}
	
	
	func getOptionResponse() map[string]interface{} {
		response := map[string]interface{}{
			"endpoints": map[string]interface{}{
				"/run-planets": map[string]interface{}{
					"parameters": map[string]interface{}{
						"birthdate":   "string (required) - Birthdate in the format of 'dd.mm.yyyy', 'dd-mm-yyyy', or 'dd/mm/yyyy'",
						"utctime":     "string (required) - UTC time in the format of 'hh:mm'",
						"latitude":    "string (required) - Latitude in decimal format",
						"longitude":   "string (required) - Longitude in decimal format",
						"altitude":    "string (required) - Altitude in meters",
						"housesystem": "string (required) - House system (P for Placidus, R for Regiomontanus)",
					},
				},
				"/run-houses": map[string]interface{}{
					"parameters": map[string]interface{}{
						"birthdate":   "string (required) - Birthdate in the format of 'dd.mm.yyyy', 'dd-mm-yyyy', or 'dd/mm/yyyy'",
						"utctime":     "string (required) - UTC time in the format of 'hh:mm'",
						"latitude":    "string (required) - Latitude in decimal format",
						"longitude":   "string (required) - Longitude in decimal format",
						"altitude":    "string (required) - Altitude in meters",
						"housesystem": "string (required) - House system (P for Placidus, R for Regiomontanus)",
					},
				},
				"/run-star": map[string]interface{}{
					"parameters": map[string]interface{}{
						"birthdate": "string (required) - Birthdate in the format of 'dd.mm.yyyy', 'dd-mm-yyyy', or 'dd/mm/yyyy'",
						"utctime":   "string (required) - UTC time in the format of 'hh:mm'",
					},
				},
			},
		}
	
		return response
	}
	