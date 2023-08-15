# daysUntil

This is a simple terminal tool for showing the time until or since a given date/time.

### Usage

Multiple dates can be passed in the command line. Acceptable formats are either:

    mm-dd-yyyy
    mm-dd-yyyy HH:mm   (24hour clock)


### Example

    daysUntil  "12-10-2021 13:00" "12-11-2021 13:00" "12-17-2021 17:00"

### Notes

- If the date is in the past, "day since" is display
- A date without a time is calculated to the start of the day
- Times are local