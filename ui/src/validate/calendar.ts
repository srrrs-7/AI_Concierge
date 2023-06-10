import { z } from 'zod';

// string validation
export const validateString = z.string().min(3).max(10);

// english validation
const englishRegex = /^[A-Z]+$/;
export const validateEnglish = z.string().regex(englishRegex);

// number validation
const numberRegex = /^[1-9]+$/;
export const validateDate = z.string().regex(numberRegex);

export const validateWeek = z.union([
  z.literal('Sunday'),
  z.literal('Monday'),
  z.literal('Tuesday'),
  z.literal('Wednesday'),
  z.literal('Thursday'),
  z.literal('Friday'),
  z.literal('Saturday')
]);

export const validateMonth = z.union([
  z.literal('January'),
  z.literal('February'),
  z.literal('March'),
  z.literal('April'),
  z.literal('May'),
  z.literal('June'),
  z.literal('July'),
  z.literal('August'),
  z.literal('September'),
  z.literal('October'),
  z.literal('November'),
  z.literal('December')
]);

// time validation 10:30:00 new Date()
export const validateTime = z.time();

// date validation yyy-mm-dd, unixtime
export const validateDay = z.date();

// date validation yyy-mm-dd
const dateRegex = /^\d{4}-\d{2}-\d{2}$/;
export const dateSchema = z.string().regex(dateRegex);

// time validation year
export const validateYear = z
  .date()
  .min(new Date('2023-01-01'))
  .max(new Date('2023-12-31'));

// each month validation
export const validateJanuaryDate = z.number().min(0).max(31);
export const validateFebruaryDate = z.number().min(0).max(29);
export const validateMarchDate = z.number().min(0).max(31);
export const validateAprilDate = z.number().min(0).max(30);
export const validateMayDate = z.number().min(0).max(31);
export const validateJuneDate = z.number().min(0).max(30);
export const validateJulyDate = z.number().min(0).max(31);
export const validateAugustDate = z.number().min(0).max(31);
export const validateSeptemberDate = z.number().min(0).max(30);
export const validateOctoberDate = z.number().min(0).max(31);
export const validateNovemberDate = z.number().min(0).max(30);
export const validateDecemberDate = z.number().min(0).max(31);
