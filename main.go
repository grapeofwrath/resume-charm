package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/glamour"
)

type Contact struct {
    Name string
    Location string
}

type Duty struct {
    Title string
    Skills []string
}

type Experience struct {
    Employer string
    Location string
    Date string
    Title string
    Duties []Duty
}

type Education struct {
    Name string
    Location string
    Program string
    Date string
}

func main() {
    m2 := Contact{
        "Marcus Montgomery",
        "Lebanon, TN",
    }

    var experience []Experience
    pcr := Experience{
        "Paramount Construction & Restoration",
        "Nashville, TN",
        "NOV 2021 - PRESENT",
        "Flood & Restoration Lead/Trim Carpenter",
        []Duty{
            {
                "Flood Lead",
                []string{
                    "Responding to emergency situations and restoring calm to the chaos",
                    "Constructing a plan of action based on the severity of each situation and the help present",
                    "Executing the plan in an efficient and timely fashion",
                    "Ensuring the client is sufficiently informed and at ease",
                },
            },
            {
                "Restoration Lead",
                []string{
                    "Leading a crew and maintaining a high standard for the work accomplished in a timely manner",
                    "Remediating any damage from water, mold, or fire",
                    "Making any repairs from floods or other incidents typically including drywall, paint, and trim work",
                    "Planning each day’s scope based on the overall scope",
                    "Occasionally conducting repairs on plumbing, electrical, and framing",
                },
            },
        },
    }
    mfrm := Experience{
        "Mattress Firm",
        "Kitsap County, WA",
        "NOV 2017 - JUN 2021",
        "Assistant Store Manager",
        []Duty{
            {
                "Assistant Store Manager",
                []string{
                    "Creating an environment where the customer is at the center by cultivating strong relationships",
                    "Driving individual and store team’s performance, through the achievement of KPI’s, to reach targeted goals",
                    "Managing assets for my assigned store",
                    "Ensuring company policies, including appearance and functionality standards, and state regulations are followed",
                },
            },
        },
    }
    experience = append(experience, pcr, mfrm)

    var education []Education
    fi := Education{
        "Flatiron School",
        "Online",
        "Cybersecurity Analytics",
        "JUL 2020 - NOV 2020",
    }
    frcs := Education{
        "Franklin Road Christian School",
        "Murfreesboro, TN",
        "High School Diploma",
        "AUG 2012 - MAY 2016",
    }
    education = append(education, fi, frcs)

    in := `# `+m2.Name+`

`+m2.Location+`

---

## EXPERIENCE

`+fmtExperience(experience)+`

## EDUCATION

`+fmtEducation(education)+`
`

    out, err := glamour.Render(in, "dark")
    if err != nil {
        log.Fatal("OPE! Ya got an error there bud:", err)
    }
    fmt.Print(out)
}

func fmtEducation(e []Education) string {
    var edus []string
    for _, edu := range e {
        ed := "### **"+edu.Name+"**, "+edu.Location+"\n\n*"+edu.Program+" - "+edu.Date+"*\n"
        edus = append(edus, ed)
    }
    edusStr := strings.Join(edus, "\n\n")
    return edusStr
}

func fmtExperience(e []Experience) string {
    var exps []string
    for _, ex := range e {
        title := "### **"+ex.Employer+"** - "+ex.Location+"\n\n*"+ex.Date+" - "+ex.Title+"*\n\n"
        var duties []string
        for _, d := range ex.Duties {
            var skills []string
            for _, s := range d.Skills {
                skill := " - "+s
                skills = append(skills, skill)
            }
            skillsStr := strings.Join(skills, "\n")
            duty := "**"+d.Title+"**\n\n"+skillsStr
            duties = append(duties, duty)
        }
        dutiesStr := strings.Join(duties, "\n\n")
        exp := title+dutiesStr
        exps = append(exps, exp)
    }
    expsStr := strings.Join(exps, "\n\n")
    return expsStr
}
